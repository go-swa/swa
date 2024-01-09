package service_config

import "github.com/ServiceWeaver/weaver"

type ModSystem struct {
	weaver.AutoMarshal
	Env string `mapstructure:"env" json:"env" yaml:"env"`
	Addr int `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	OssType string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`
	UseMultipoint bool `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`
	UseRedis bool `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`
	LimitCountIP int `mapstructure:"ip-limit-count" json:"ip-limit-count" yaml:"ip-limit-count"`
	LimitTimeIP int `mapstructure:"ip-limit-time" json:"ip-limit-time" yaml:"ip-limit-time"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
