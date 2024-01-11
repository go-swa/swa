package service_config

import "github.com/ServiceWeaver/weaver"

type ModRedis struct {
	weaver.AutoMarshal
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
