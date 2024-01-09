package service_config

import (
	"context"
	"demo1/utils"
)

func (imp *impl) GetCfgSwaConfig(ctx context.Context) (ModSwaConfig, error) {
	imp.swaConfig.GormDB.SetDsn()
	return imp.swaConfig, nil
}

func (imp *impl) SetCfgGormDb(_ context.Context, m ModGormDb) error {
	imp.swaConfig.GormDB = m
	return nil
}

func (imp *impl) GetCfgGormDb(ctx context.Context) (ModGormDb, error) {
	imp.swaConfig.GormDB.SetDsn()
	return imp.swaConfig.GormDB, nil
}

func (imp *impl) GetCfgJwt(ctx context.Context) (ModJwt, error) {
	return imp.swaConfig.Jwt, nil
}


func (imp *impl) SetSystemConfig(_ context.Context, system ModSwaConfig) (err error) {
	cs := utils.StructToMap(system)
	for k, v := range cs {
		imp.swaViper.Set(k, v)
	}
	err = imp.swaViper.WriteConfig()
	return err
}
