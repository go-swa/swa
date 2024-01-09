package service_swa

import (
	"context"
	"demo1/utils"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type TUser interface {
	Login(ctx context.Context, m SysUser) (SysUser, error)
	GetUserInfo(ctx context.Context, uuid uuid.UUID) (SysUser, error)
	Register(ctx context.Context, m SysUser) (SysUser, error)
	ChangePassword(ctx context.Context, uid uint, oldPassword string, newPassword string) (userInter *SysUser, err error)
	GetUserInfoList(ctx context.Context, page PageInfo) (list []SysUser, total int64, err error)
	SetUserRole(_ context.Context, id uint, roleId uint) (err error)
	SetUserAuthorities(_ context.Context, id uint, roleIds []uint) (err error)
	SetUserInfo(_ context.Context, req SysUser) error
	DeleteUser(_ context.Context, id int) (err error)
	SetSelfInfo(_ context.Context, req SysUser) error
	ResetPassword(_ context.Context, ID uint) (err error)
	GetUserDict(_ context.Context) (user []SysUser, err error)


}


func (imp *impl) Login(ctx context.Context, login SysUser) (user SysUser, err error) {
	if nil == imp.gormDB {
		return user, fmt.Errorf("db not init")
	}

	err = imp.gormDB.Where("username = ?", login.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(login.Password, user.Password); !ok {

			return user, errors.New("密码错误")
		}

		var role SwaRole
		err = imp.gormDB.Where("roleId = ?", user.RoleId).First(&role).Error
		if err != nil {
			return user, err
		}
		user.SwaRole = role

		var roles []SwaRole
		err = imp.gormDB.Where("roleId = ?", user.RoleId).Find(&roles).Error
		if err != nil {
			return user, err
		}
		user.Authorities = roles

		imp.UserRoleDefaultRouter(&user)
	}
	if err != nil {
		return user, err
	}
	if user.Enable != 1 {
		return user, fmt.Errorf("用户被禁止登录")
	}

	return user, err
}

func (imp *impl) GetUserInfo(ctx context.Context, uuid uuid.UUID) (user SysUser, err error) {
	var modUser SysUser
	var role SwaRole
	var roles []SwaRole

	err = imp.gormDB.First(&modUser, "uuid = ?", uuid).Error
	if err != nil {
		return modUser, err
	}

	err = imp.gormDB.First(&role, "roleId = ?", modUser.RoleId).Error
	if err != nil {
		return user, err
	} else {
		modUser.SwaRole = role
	}

	err = imp.gormDB.Find(&roles, "roleId = ?", modUser.RoleId).Error
	if err != nil {
		return user, err
	} else {
		modUser.Authorities = roles
	}

	imp.UserRoleDefaultRouter(&modUser)

	return modUser, err
}

func (imp *impl) UserRoleDefaultRouter(user *SysUser) {
	var menuIds []string
	err := imp.gormDB.Model(&Role2Menus{}).Where("role_id = ?", user.RoleId).Pluck("menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am SysBaseMenu
	err = imp.gormDB.First(&am, "name = ? and id in (?)", user.SwaRole.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.SwaRole.DefaultRouter = "404"
	}
}


func (imp *impl) Register(_ context.Context, m SysUser) (SysUser, error) {
	var authorities []SwaRole
	var userInter SysUser
	var err error

	user := &SysUser{
		Username:    m.Username,
		NickName:    m.NickName,
		Password:    m.Password,
		HeaderImg:   m.HeaderImg,
		RoleId:      m.RoleId,
		Authorities: authorities,
		Enable:      m.Enable,
		Phone:       m.Phone,
		Email:       m.Email,
	}

	if !errors.Is(imp.gormDB.Where("username = ?", m.Username).First(&user).Error, gorm.ErrRecordNotFound) { // --判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	m.Password = utils.BcryptHash(m.Password)
	m.UUID = uuid.Must(uuid.NewV4())
	err = imp.gormDB.Create(&m).Error

	return m, err
}


func (imp *impl) ChangePassword(_ context.Context, uid uint, oldPassword string, newPassword string) (userInter *SysUser, err error) {
	var user SysUser
	if err = imp.gormDB.Where("id = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(oldPassword, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = imp.gormDB.Save(&user).Error
	return &user, err

}


func (imp *impl) GetUserInfoList(_ context.Context, pageInfo PageInfo) (list []SysUser, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := imp.gormDB.Model(&SysUser{})
	var userList []SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error

	for i, v := range userList {
		var roleIds []uint
		err := imp.gormDB.Model(&UserRole{}).Where("user_id = ?", v.ID).Pluck("role_id", &roleIds).Error
		if err != nil {
			return userList, total, errors.New("查询用户角色id失败")
		}
		userList[i].RoleIds = roleIds

		var roles []SwaRole
		err = imp.gormDB.Where("roleId in (?)", roleIds).Find(&roles).Error
		if err != nil {
			return userList, total, errors.New("查询用户角色内容失败")
		}
		userList[i].Authorities = roles
	}
	return userList, total, err
}


func (imp *impl) SetUserRole(_ context.Context, id uint, roleId uint) (err error) {
	assignErr := imp.gormDB.Where("user_id = ? AND role_id = ?", id, roleId).First(&UserRole{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = imp.gormDB.Where("id = ?", id).First(&SysUser{}).Update("role_id", roleId).Error
	return err
}


func (imp *impl) SetUserAuthorities(_ context.Context, id uint, roleIds []uint) (err error) {
	return imp.gormDB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]UserRole{}, "user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useRole []UserRole
		for _, v := range roleIds {
			useRole = append(useRole, UserRole{
				UserId: id, RoleId: v,
			})
		}
		TxErr = tx.Create(&useRole).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&SysUser{}).Update("role_id", roleIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}


func (imp *impl) DeleteUser(_ context.Context, id int) (err error) {
	var user SysUser
	err = imp.gormDB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	err = imp.gormDB.Delete(&[]UserRole{}, "user_id = ?", id).Error
	return err
}


func (imp *impl) SetUserInfo(_ context.Context, req SysUser) error {
	return imp.gormDB.Model(&SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"side_mode":  req.SideMode,
			"enable":     req.Enable,
		}).Error
}


func (imp *impl) SetSelfInfo(_ context.Context, req SysUser) error {
	return imp.gormDB.Model(&SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}




func (imp *impl) FindUserById(id int) (user *SysUser, err error) {
	var u SysUser
	err = imp.gormDB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}


func (imp *impl) FindUserByUuid(uuid string) (user *SysUser, err error) {
	var u SysUser
	if err = imp.gormDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}


func (imp *impl) ResetPassword(_ context.Context, ID uint) (err error) {
	err = imp.gormDB.Model(&SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}


func (imp *impl) GetUserDict(_ context.Context) (user []SysUser, err error) {
	err = imp.gormDB.Find(&user).Select("ID,username,nick_name").Error
	return user, err

}
