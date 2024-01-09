package service_config

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type T interface {
	SetCfgGormDb(ctx context.Context, m ModGormDb) error

	GetCfgSwaConfig(ctx context.Context) (ModSwaConfig, error)
	GetCfgGormDb(ctx context.Context) (ModGormDb, error)
	GetCfgJwt(ctx context.Context) (ModJwt, error)
	SetSystemConfig(_ context.Context, system ModSwaConfig) (err error)
}

type impl struct {
	weaver.Implements[T]
	swaConfig ModSwaConfig
	swaViper  *viper.Viper
}

func (imp *impl) Init(ctx context.Context) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	v.AddConfigPath("./service_config")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取swa toml配置文件失败: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		imp.Logger(ctx).Info("config file changed:")
		if err = v.Unmarshal(&imp.swaConfig); err != nil {
			imp.Logger(ctx).Info("系统配置数据被修改", zap.Error(err))
		}
	})

	if err = v.Unmarshal(&imp.swaConfig); err != nil {
		imp.Logger(ctx).Info("系统配置数据转换失败", zap.Error(err))
	}
	imp.swaViper = v
	return nil
}
