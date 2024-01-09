package service_file

import (
	"context"
	"demo1/service/service_config"
	"github.com/ServiceWeaver/weaver"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type T interface {
	TUpDownload
}

type impl struct {
	weaver.Implements[T]
	gormDB        *gorm.DB
	serviceConfig weaver.Ref[service_config.T]
	swaConfig     service_config.ModSwaConfig
}

func (imp *impl) Init(ctx context.Context) (err error) {
	imp.swaConfig, err = imp.serviceConfig.Get().GetCfgSwaConfig(ctx)
	if err != nil {
		panic(err)
	}
	db, err := getGormDB("mysql", "root:admin@tcp(192.168.3.61:3306)/swa-demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		imp.Logger(ctx).Error("get gorm db fail:%v", err)
		panic(err)
	}
	imp.gormDB = db

	err = imp.gormDB.AutoMigrate(
		ModUpload{},
	)
	return err
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
