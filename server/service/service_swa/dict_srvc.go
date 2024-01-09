package service_swa

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type TDict interface {
	CreateSwaDict(_ context.Context, swaDict ModDict) (err error)
	DeleteSwaDict(_ context.Context, swaDict ModDict) (err error)
	UpdateSwaDict(_ context.Context, swaDict ModDict) (err error)
	GetSwaDict(_ context.Context, Type string, Id uint, status *bool) (swaDict ModDict, err error)
	GetSwaDictInfoList(_ context.Context, info SwaDictSearch) (list []ModDict, total int64, err error)
}

func (imp *impl) CreateSwaDict(_ context.Context, swaDict ModDict) (err error) {
	if (!errors.Is(imp.gormDB.First(&ModDict{}, "type = ?", swaDict.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = imp.gormDB.Create(&swaDict).Error
	return err
}


func (imp *impl) DeleteSwaDict(_ context.Context, swaDict ModDict) (err error) {
	err = imp.gormDB.Where("id = ?", swaDict.ID).Preload("SwaDictails").First(&swaDict).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = imp.gormDB.Delete(&swaDict).Error
	if err != nil {
		return err
	}

	if swaDict.ModDictails != nil {
		return imp.gormDB.Where("swaDictID=?", swaDict.ID).Delete(swaDict.ModDictails).Error
	}
	return
}


func (imp *impl) UpdateSwaDict(_ context.Context, swaDict ModDict) (err error) {
	var dict ModDict
	sysDictionaryMap := map[string]interface{}{
		"Name":   swaDict.Name,
		"Type":   swaDict.Type,
		"Status": swaDict.Status,
		"Detail": swaDict.Detail,
	}
	db := imp.gormDB.Where("id = ?", swaDict.ID).First(&dict)
	if dict.Type != swaDict.Type {
		if !errors.Is(imp.gormDB.First(&ModDict{}, "type = ?", swaDict.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}
	err = db.Updates(sysDictionaryMap).Error
	return err
}


func (imp *impl) GetSwaDict(_ context.Context, Type string, Id uint, status *bool) (swaDict ModDict, err error) {
	var flag = false
	if status == nil {
		flag = true
	} else {
		flag = *status
	}
	err = imp.gormDB.Where("(type = ? OR id = ?) and status = ?", Type, Id, flag).First(&swaDict).Error
	if err != nil {
		return swaDict, err
	}
	var dictail []ModDictail
	err = imp.gormDB.Where("swaDictID = ? and status = ?", swaDict.ID, true).Order("sort").First(&dictail).Error
	if err != nil {
		return swaDict, err
	}
	swaDict.ModDictails = dictail
	return swaDict, err
}


func (imp *impl) GetSwaDictInfoList(_ context.Context, info SwaDictSearch) (list []ModDict, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := imp.gormDB.Model(&ModDict{})
	var sysDictionarys []ModDict
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Detail != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Detail+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return sysDictionarys, total, err
}
