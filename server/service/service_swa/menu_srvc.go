package service_swa

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TMenu interface {


	GetRoleMenu(ctx context.Context, roleId uint) ([]SysBaseMenu, error)
	GetInfoList(ctx context.Context) (list []SysBaseMenu, total int64, err error)
	AddMenuRole(ctx context.Context, menus []SysBaseMenu, roleId uint) (err error)
	GetBaseMenuById(_ context.Context, id int) (menu SysBaseMenu, err error)
	GetAllBaseMenu(_ context.Context) (menus []SysBaseMenu, err error)
	AddBaseMenu(_ context.Context, menu SysBaseMenu) error
	DeleteBaseMenu(_ context.Context, id int) (err error)
	UpdateBaseMenu(ctx context.Context, menu SysBaseMenu) (err error)
	GetMenuRole(_ context.Context, id uint) (menus []SysBaseMenu, err error)
}

func (imp *impl) GetRoleMenu(_ context.Context, roleId uint) (menus []SysBaseMenu, err error) {
	var MenuIds []string
	err = imp.gormDB.Model(&Role2Menus{}).Where("role_id = ?", roleId).Pluck("menu_id", &MenuIds).Error
	if err != nil {
		return
	}


	err = imp.gormDB.Where("id in (?)", MenuIds).Order("sort").Find(&menus).Error
	if err != nil {
		return menus, err
	}

	return menus, err
}







func (imp *impl) GetInfoList(_ context.Context) (list []SysBaseMenu, total int64, err error) {
	var allMenus []SysBaseMenu
	err = imp.gormDB.Order("sort").Find(&allMenus).Error
	return allMenus, total, err
}




func (imp *impl) AddBaseMenu(_ context.Context, menu SysBaseMenu) error {
	if !errors.Is(imp.gormDB.Where("name = ?", menu.Name).First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return imp.gormDB.Create(&menu).Error
}


func (imp *impl) getBaseMenuTreeMap() (treeMap map[string][]SysBaseMenu, err error) {
	var allMenus []SysBaseMenu
	treeMap = make(map[string][]SysBaseMenu)
	err = imp.gormDB.Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}




func (imp *impl) AddMenuRole(ctx context.Context, menus []SysBaseMenu, roleId uint) (err error) {
	var auth RoleUserMenu
	auth.RoleId = roleId
	auth.SysBaseMenus = menus
	return err
}


func (imp *impl) GetMenuRole(_ context.Context, id uint) (menus []SysBaseMenu, err error) {
	var baseMenu []SysBaseMenu
	var roleMenu []Role2Menus
	err = imp.gormDB.Where("role_id = ?", id).Find(&roleMenu).Error

	if err != nil {
		return
	}

	var MenuIds []uint

	for i := range roleMenu {
		MenuIds = append(MenuIds, roleMenu[i].MenuId)
	}
	err = imp.gormDB.Where("id in (?) ", MenuIds).Order("sort").Find(&baseMenu).Error

	return baseMenu, err
}




func (imp *impl) GetBaseMenuById(_ context.Context, id int) (menu SysBaseMenu, err error) {
	err = imp.gormDB.Where("id = ?", id).First(&menu).Error
	return
}

func (imp *impl) GetAllBaseMenu(_ context.Context) (menus []SysBaseMenu, err error) {
	err = imp.gormDB.Order("sort").Find(&menus).Error
	return
}


func (imp *impl) DeleteBaseMenu(_ context.Context, id int) (err error) {
	err = imp.gormDB.Where("parent_id = ?", id).First(&SysBaseMenu{}).Error
	if err != nil {
		err := imp.gormDB.Delete(&SwaRoleMenu{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			if fmt.Sprintf("%v", err) != "record not found" {
				return errors.New("删除菜单角色数据错误")
			}
		}
		err = imp.gormDB.Delete(&SysBaseMenu{}, "id = ?", id).Error
		if err != nil {
			if fmt.Sprintf("%v", err) != "record not found" {
				return errors.New("删除菜单数据错误")
			}
		}
		err = imp.gormDB.Delete(&SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			if fmt.Sprintf("%v", err) != "record not found" {
				return errors.New("删除菜单参数数据错误")
			}
		}
		err = imp.gormDB.Delete(&SysBaseMenuBtn{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			if fmt.Sprintf("%v", err) != "record not found" {
				return errors.New("删除菜单按键数据错误")
			}
		}
		err = imp.gormDB.Delete(&SwaRoleBtn{}, "sys_menu_id = ?", id).Error
		if err != nil {
			if fmt.Sprintf("%v", err) != "record not found" {
				return errors.New("删除角色按键数据错误")
			}
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return nil
}


func (imp *impl) UpdateBaseMenu(ctx context.Context, menu SysBaseMenu) (err error) {
	var oldMenu SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["active_name"] = menu.ActiveName
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = imp.gormDB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				imp.Logger(ctx).Debug("存在相同name修改失败")
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			imp.Logger(ctx).Debug(txErr.Error())
			return txErr
		}
		txErr = tx.Unscoped().Delete(&SysBaseMenuBtn{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			imp.Logger(ctx).Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				imp.Logger(ctx).Debug(txErr.Error())
				return txErr
			}
		}

		if len(menu.MenuBtn) > 0 {
			for k := range menu.MenuBtn {
				menu.MenuBtn[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.MenuBtn).Error
			if txErr != nil {
				imp.Logger(ctx).Debug(txErr.Error())
				return txErr
			}
		}

		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			imp.Logger(ctx).Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	return err
}
