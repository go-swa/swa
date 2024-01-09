package service_config

import (
	"github.com/ServiceWeaver/weaver"
)

type ModSwaConfig struct {
	weaver.AutoMarshal
	Jwt   ModJwt   `mapstructure:"jwt" json:"jwt" toml:"jwt"`
	Redis ModRedis `mapstructure:"redis" json:"redis" toml:"redis"`
	Captcha ModCaptcha `mapstructure:"captcha" json:"captcha" toml:"captcha"`
	GormDB  ModGormDb  `mapstructure:"gormDB" json:"gormDB" toml:"gormDB"`
	System ModSystem `mapstructure:"system" json:"system" toml:"system"`
	Local  Local     `mapstructure:"local" json:"local" toml:"local"`
}
