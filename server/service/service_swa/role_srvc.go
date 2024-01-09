package service_swa

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type TRole interface {



	CreateRole(_ context.Context, auth SwaRole) (swaRole SwaRole, err error)

	UpdateRole(_ context.Context, auth SwaRole) (swaRole SwaRole, err error)
	DeleteRole(_ context.Context, auth RoleUserMenu) (err error)
	GetRoleList(_ context.Context, info PageInfo) (list []SwaRole, total int64, err error)
	GetRoleInfo(_ context.Context, id uint) (sa SwaRole, err error)
	SetSubRoles(_ context.Context, auth SwaRole) error
	SetRoleMenus(_ context.Context, auth SwaRole) error
	GetSubRoles(_ context.Context, id uint) (subRoles []SwaRole, err error)
}

var ErrRoleExistence = errors.New("存在相同角色id")

func (imp *impl) CreateRole(_ context.Context, auth SwaRole) (swaRole SwaRole, err error) {
	var roleBox SwaRole
	if !errors.Is(imp.gormDB.Where("roleId = ?", auth.RoleId).First(&roleBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}
	err = imp.gormDB.Create(&auth).Error
	return auth, err
}




func (imp *impl) UpdateRole(_ context.Context, auth SwaRole) (swaRole SwaRole, err error) {
	err = imp.gormDB.Where("role_id = ?", auth.RoleId).First(&SwaRole{}).Updates(&auth).Error
	return auth, err
}


func (imp *impl) DeleteRole(_ context.Context, auth RoleUserMenu) (err error) {
	if errors.Is(imp.gormDB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(imp.gormDB.Where("roleId = ?", auth.RoleId).First(&SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(imp.gormDB.Where("parentId = ?", auth.RoleId).First(&SwaRole{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := imp.gormDB.Preload("SysBaseMenus").Preload("DataRoleId").Where("roleId = ?", auth.RoleId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		err = imp.gormDB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
	}
	err = imp.gormDB.Delete(&[]UserRole{}, "role_id = ?", auth.RoleId).Error
	if err != nil {
		return
	}
	err = imp.gormDB.Delete(&[]SwaRoleBtn{}, "role_id = ?", auth.RoleId).Error
	if err != nil {
		return
	}
	roleId := strconv.Itoa(int(auth.RoleId))
	imp.ClearCasbin(0, roleId)
	return err
}


func (imp *impl) GetRoleList(_ context.Context, info PageInfo) (list []SwaRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := imp.gormDB.Model(&SwaRole{})
	if err = db.Where("parentId = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var swaRole []SwaRole
	err = db.Limit(limit).Offset(offset).Where("parentId = ?", "0").Find(&swaRole).Error
	return swaRole, total, err
}


func (imp *impl) GetRoleInfo(_ context.Context, id uint) (sa SwaRole, err error) {
	err = imp.gormDB.Preload("DataRoleId").Where("roleId = ?", id).First(&sa).Error
	return sa, err
}


func (imp *impl) SetSubRoles(_ context.Context, swaRole SwaRole) error {
	var oldSubRoles []Role2SubRoles
	err := imp.gormDB.Where("role_id = ?", swaRole.RoleId).Find(&oldSubRoles).Error
	if err != nil {
		return err
	}
	if len(swaRole.SubRoles) == 0 {
		for _, oldSubRole := range oldSubRoles {
			err = imp.gormDB.Where("role_id = ? and sub_role_id =?", oldSubRole.RoleId, oldSubRole.SubRoleId).Delete(oldSubRole).Error
			if err != nil {
				return err
			}
		}
		return nil
	}
	if len(oldSubRoles) < 1 {
		for _, newSubRole := range swaRole.SubRoles {
			var oneSubRole = Role2SubRoles{RoleId: swaRole.RoleId, SubRoleId: newSubRole.RoleId}
			err = imp.gormDB.Create(oneSubRole).Error
			if err != nil {
				return err
			}
		}
		return nil
	}
	for _, oldSubRole := range oldSubRoles {
		for idxNew, newSubRole := range swaRole.SubRoles {
			if oldSubRole.RoleId == swaRole.RoleId && oldSubRole.SubRoleId == newSubRole.RoleId {
				break
			}
			if idxNew == len(swaRole.SubRoles)-1 {
				err = imp.gormDB.Where("role_id = ? and sub_role_id =?", oldSubRole.RoleId, oldSubRole.SubRoleId).Delete(oldSubRole).Error
				if err != nil {
					return err
				}
			}

		}
	}

	for _, newSubRole := range swaRole.SubRoles {
		for idxOld, oldSubRole := range oldSubRoles {
			if oldSubRole.RoleId == swaRole.RoleId && oldSubRole.SubRoleId == newSubRole.RoleId {
				break
			}
			if idxOld == len(oldSubRoles)-1 {
				var oneSubRole = Role2SubRoles{RoleId: swaRole.RoleId, SubRoleId: newSubRole.RoleId}
				err = imp.gormDB.Create(oneSubRole).Error
				if err != nil {
					return err
				}
			}

		}
	}

	return err
}

func (imp *impl) SetRoleMenus(_ context.Context, auth SwaRole) error {
	var oldRoleMenus []Role2Menus
	err := imp.gormDB.Where("role_id = ?", auth.RoleId).Find(&oldRoleMenus).Error
	if err != nil {
		return err
	}

	if len(auth.RoleMenus) == 0 {
		for _, oldRoleMenu := range oldRoleMenus {
			err = imp.gormDB.Where("role_id = ? and menu_id =?", oldRoleMenu.RoleId, oldRoleMenu.MenuId).Delete(oldRoleMenu).Error
			if err != nil {
				return err
			}

		}
		return nil
	}
	if len(oldRoleMenus) < 1 {
		for _, oneNewMenu := range auth.RoleMenus {
			var oneNewRoleMenu = Role2Menus{RoleId: auth.RoleId, MenuId: oneNewMenu.ID}
			err = imp.gormDB.Create(oneNewRoleMenu).Error
			if err != nil {
				return err
			}
		}
		return nil
	}
	for _, oldRoleMenu := range oldRoleMenus {
		for idxNew, newMenu := range auth.RoleMenus {
			if oldRoleMenu.RoleId == auth.RoleId && oldRoleMenu.MenuId == newMenu.ID {
				break
			}
			if idxNew == len(auth.RoleMenus)-1 {
				err = imp.gormDB.Where("role_id = ? and menu_id =?", oldRoleMenu.RoleId, oldRoleMenu.MenuId).Delete(oldRoleMenu).Error
				if err != nil {
					return err
				}
			}

		}
	}

	for _, oneNewMenu := range auth.RoleMenus {
		for idxOld, oldRoleMenu := range oldRoleMenus {
			if oldRoleMenu.RoleId == auth.RoleId && oldRoleMenu.MenuId == oneNewMenu.ID {
				break
			}
			if idxOld == len(oldRoleMenus)-1 {
				var oneNewRoleMenu = Role2Menus{RoleId: auth.RoleId, MenuId: oneNewMenu.ID}
				err = imp.gormDB.Create(oneNewRoleMenu).Error
				if err != nil {
					return err
				}
			}

		}
	}
	return err
}



func (imp *impl) GetSubRoles(_ context.Context, id uint) (subRoles []SwaRole, err error) {
	var subRoleIds []string
	err = imp.gormDB.Model(&Role2SubRoles{}).Where("role_id = ?", id).Pluck("sub_role_id", &subRoleIds).Error
	if err != nil {
		return
	}

	err = imp.gormDB.Where("roleId in (?)", subRoleIds).Order("roleId").Find(&subRoles).Error
	if err != nil {
		return
	}

	return subRoles, err
}
