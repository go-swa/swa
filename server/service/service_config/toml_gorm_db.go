package service_config

import (
	"fmt"
	"github.com/ServiceWeaver/weaver"
)

type ModGormDb struct {
	weaver.AutoMarshal
	DbType       string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`
	Ip           string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`
	Dsn          string
}

func (m *ModGormDb) SetDsn() {
	m.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		m.Username,
		m.Password,
		m.Ip,
		m.Port,
		m.Dbname,
		m.Config,
	)
}
