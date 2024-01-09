package service_swa

import (
	"context"
)

type TRecord interface {
	CreateSwaRecord(_ context.Context, swaRecord SwaRecord) (err error)
	DeleteSwaRecordByIds(_ context.Context, ids []uint) (err error)
	DeleteSwaRecord(_ context.Context, swaRecord SwaRecord) (err error)
	GetSwaRecord(_ context.Context, id uint) (swaRecord SwaRecord, err error)
	GetSwaRecordInfoList(_ context.Context, info SwaRecordSearch) (list []SwaRecord, total int64, err error)
}

func (imp *impl) CreateSwaRecord(_ context.Context, swaRecord SwaRecord) (err error) {
	if len(swaRecord.Body) > 1000 {
		swaRecord.Body = swaRecord.Body[0:999]
	}
	err = imp.gormDB.Create(&swaRecord).Error
	return err
}


func (imp *impl) DeleteSwaRecordByIds(_ context.Context, ids []uint) (err error) {
	err = imp.gormDB.Delete(&[]SwaRecord{}, "id in (?)", ids).Error
	return err
}


func (imp *impl) DeleteSwaRecord(_ context.Context, swaRecord SwaRecord) (err error) {
	err = imp.gormDB.Delete(&swaRecord).Error
	return err
}


func (imp *impl) GetSwaRecord(_ context.Context, id uint) (swaRecord SwaRecord, err error) {
	err = imp.gormDB.Where("id = ?", id).First(&swaRecord).Error
	return
}


func (imp *impl) GetSwaRecordInfoList(_ context.Context, info SwaRecordSearch) (list []SwaRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := imp.gormDB.Model(&SwaRecord{})
	var swaRecords []SwaRecord
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&swaRecords).Error
	return swaRecords, total, err
}
