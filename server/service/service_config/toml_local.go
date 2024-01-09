package service_config

import "github.com/ServiceWeaver/weaver"

type Local struct {
	weaver.AutoMarshal
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"`
}
