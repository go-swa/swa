package service_file

import (
	"context"
)

type TUpDownload interface {
	Upload(_ context.Context, file ModUpload) error
	FindFile(_ context.Context, id uint) (ModUpload, error)
	DeleteFile(ctx context.Context, file ModUpload) (err error)
	EditFileName(_ context.Context, file ModUpload) (err error)
	GetFileRecordInfoList(_ context.Context, info PageInfo) (list []ModUpload, total int64, err error)
}


func (imp *impl) Upload(_ context.Context, file ModUpload) error {
	return imp.gormDB.Create(&file).Error
}


func (imp *impl) FindFile(_ context.Context, id uint) (ModUpload, error) {
	var file ModUpload
	err := imp.gormDB.Where("id = ?", id).First(&file).Error
	return file, err
}


func (imp *impl) DeleteFile(_ context.Context, file ModUpload) (err error) {
	err = imp.gormDB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

func (imp *impl) EditFileName(_ context.Context, file ModUpload) (err error) {
	var fileFromDb ModUpload
	return imp.gormDB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}


func (imp *impl) GetFileRecordInfoList(_ context.Context, info PageInfo) (list []ModUpload, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := imp.gormDB.Model(&ModUpload{})
	var fileLists []ModUpload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}


