package service_swa

import (
	"context"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"strconv"
	"sync"
)

type TCasbin interface {


	UpdateCasbin(_ context.Context, RoleID uint, casbinInfos []CasbinInfo) error
	GetPolicyPathByRoleId(_ context.Context, RoleID uint) (pathMaps []CasbinInfo, err error)
	Scasbin(_ context.Context, sub string, obj string, act string) (sok bool, err error)
}


type CasbinService struct{}


func (imp *impl) UpdateCasbin(_ context.Context, RoleID uint, casbinInfos []CasbinInfo) error {
	roleId := strconv.Itoa(int(RoleID))
	imp.ClearCasbin(0, roleId)
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{roleId, v.Path, v.Method})
	}
	e := imp.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	err := e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}


func (imp *impl) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := imp.gormDB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := imp.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return err
}


func (imp *impl) GetPolicyPathByRoleId(_ context.Context, RoleID uint) (pathMaps []CasbinInfo, err error) {
	e := imp.Casbin()
	roleId := strconv.Itoa(int(RoleID))
	list := e.GetFilteredPolicy(0, roleId)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps, nil
}


func (imp *impl) ClearCasbin(v int, p ...string) bool {
	e := imp.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}


var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (imp *impl) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(imp.gormDB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}

func (imp *impl) Scasbin(_ context.Context, sub string, obj string, act string) (sok bool, err error) {
	a, err := gormadapter.NewAdapterByDB(imp.gormDB)
	if err != nil {
		zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
		return
	}
	text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	m, err := model.NewModelFromString(text)
	if err != nil {
		zap.L().Error("字符串加载模型失败!", zap.Error(err))
		return
	}
	syncedCachedEnforcer, err = casbin.NewSyncedCachedEnforcer(m, a)
	if err != nil {
		zap.L().Error("字符串加载模型失败!", zap.Error(err))
	}
	syncedCachedEnforcer.SetExpireTime(60 * 60)
	err = syncedCachedEnforcer.LoadPolicy()
	if err != nil {
		zap.L().Error("字符串加载模型失败!", zap.Error(err))
	}
	sok, err = syncedCachedEnforcer.Enforce(sub, obj, act)
	if err != nil {
		zap.L().Error("字符串加载模型失败!", zap.Error(err))
	}
	return sok, err
}
