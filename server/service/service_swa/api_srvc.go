package service_swa

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type TApi interface {
	CreateApi(_ context.Context, api ModApi) (err error)
	DeleteApi(_ context.Context, api ModApi) (err error)
	GetAPIInfoList(_ context.Context, api ModApi, info PageInfo, order string, desc bool) (list []ModApi, total int64, err error)
	GetAllApis(_ context.Context) (apis []ModApi, err error)
	GetApiById(_ context.Context, id int) (api ModApi, err error)
	UpdateApi(_ context.Context, api ModApi) (err error)
	DeleteApisByIds(_ context.Context, ids []uint) (err error)
}

func (imp *impl) CreateApi(_ context.Context, api ModApi) (err error) {
	if !errors.Is(imp.gormDB.Where("path = ? AND method = ?", api.Path, api.Method).First(&ModApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return imp.gormDB.Create(&api).Error
}


func (imp *impl) DeleteApi(_ context.Context, api ModApi) (err error) {
	var entity ModApi
	err = imp.gormDB.Where("id = ?", api.ID).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = imp.gormDB.Delete(&entity).Error
	if err != nil {
		return err
	}
	imp.ClearCasbin(1, entity.Path, entity.Method)
	e := imp.Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}


func (imp *impl) GetAPIInfoList(_ context.Context, api ModApi, info PageInfo, order string, desc bool) (list []ModApi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := imp.gormDB.Model(&ModApi{})
	var apiList []ModApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else {
				err = fmt.Errorf("非法的排序字段: %v", order)
				return apiList, total, err
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return apiList, total, err
}


func (imp *impl) GetAllApis(_ context.Context) (apis []ModApi, err error) {
	err = imp.gormDB.Find(&apis).Error
	return
}


func (imp *impl) GetApiById(_ context.Context, id int) (api ModApi, err error) {
	err = imp.gormDB.Where("id = ?", id).First(&api).Error
	return
}


func (imp *impl) UpdateApi(_ context.Context, api ModApi) (err error) {
	var oldA ModApi
	err = imp.gormDB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(imp.gormDB.Where("path = ? AND method = ?", api.Path, api.Method).First(&ModApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = imp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = imp.gormDB.Save(&api).Error
		}
	}
	return err
}


func (imp *impl) DeleteApisByIds(_ context.Context, ids []uint) (err error) {
	var apis []ModApi
	err = imp.gormDB.Find(&apis, "id in ?", ids).Delete(&apis).Error
	if err != nil {
		return err
	} else {
		for _, sysApi := range apis {
			imp.ClearCasbin(1, sysApi.Path, sysApi.Method)
		}
		e := imp.Casbin()
		err = e.InvalidateCache()
		if err != nil {
			return err
		}
	}
	return err
}
