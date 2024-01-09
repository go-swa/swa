package service_swa

import (
	"context"
	"demo1/service/service_config"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	casbin "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type T interface {
	TCasbin
	TRole
	TUser
	TMenu
	TApi
	TDict
	TDictail
	TRecord
	TBlacklist
	TSystem
}

type impl struct {
	weaver.Implements[T]

	gormDB *gorm.DB

	serviceConfig weaver.Ref[service_config.T]
	swaConfig     service_config.ModSwaConfig
}

func (imp *impl) Init(ctx context.Context) (err error) {
	imp.swaConfig, err = imp.serviceConfig.Get().GetCfgSwaConfig(ctx)
	if err != nil {
		panic(err)
	}

	g := imp.swaConfig.GormDB
	db, err := getGormDB(g.DbType, g.Dsn)
	if err != nil {
		panic(err)
	}
	if db == nil {
		imp.Logger(ctx).Error("get swa gorm db fail")
		panic("get swa gorm db fail")
	}
	imp.gormDB = db

	fmt.Println("service_swa init:migrate...")
	err = imp.gormDB.AutoMigrate(
		ModApi{},
		SysUser{},
		SysBaseMenu{},
		JwtBlacklist{},
		SwaRole{},
		ModDict{},
		SwaRecord{},
		ModDictail{},
		SysBaseMenuParameter{},
		SwaRoleMenu{},
		UserRole{},
		SysBaseMenuBtn{},
		SwaRoleBtn{},
		casbin.CasbinRule{},

		Role2SubRoles{},
		Role2Menus{},
		Role2Apis{},
		Role2Users{},
	)
	return nil
}

func getGormDB(dbType string, dsn string) (*gorm.DB, error) {
	var err error

	if dbType == "" || dsn == "" {
		return nil, errors.Errorf("dbType为空 或 dsn为空")
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // 账号密码地址端口
	}), &gorm.Config{})

	if err != nil {
		return db, errors.Errorf("创建数据库连接失败:%v", err)
	}


	return db, err
}
