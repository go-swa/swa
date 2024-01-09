package service_swa

import "context"

type TDictail interface {
	CreateSwaDictail(_ context.Context, swaDictail ModDictail) (err error)
	DeleteSwaDictail(_ context.Context, swaDictail ModDictail) (err error)
	UpdateSwaDictail(_ context.Context, swaDictail ModDictail) (err error)
	GetSwaDictail(_ context.Context, id uint) (swaDictail ModDictail, err error)
	GetSwaDictailInfoList(_ context.Context, info SwaDictailSearch) (list []ModDictail, total int64, err error)
}

func (imp *impl) CreateSwaDictail(_ context.Context, swaDictail ModDictail) (err error) {
	err = imp.gormDB.Create(&swaDictail).Error
	return err
}


func (imp *impl) DeleteSwaDictail(_ context.Context, swaDictail ModDictail) (err error) {
	err = imp.gormDB.Delete(&swaDictail).Error
	return err
}


func (imp *impl) UpdateSwaDictail(_ context.Context, swaDictail ModDictail) (err error) {
	err = imp.gormDB.Save(swaDictail).Error
	return err
}


func (imp *impl) GetSwaDictail(_ context.Context, id uint) (swaDictail ModDictail, err error) {
	err = imp.gormDB.Where("id = ?", id).First(&swaDictail).Error
	return
}


func (imp *impl) GetSwaDictailInfoList(_ context.Context, info SwaDictailSearch) (list []ModDictail, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := imp.gormDB.Model(&ModDictail{})
	var swaDictail []ModDictail
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SwaDictID != 0 {
		db = db.Where("swaDictID = ?", info.SwaDictID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&swaDictail).Error
	return swaDictail, total, err
}
