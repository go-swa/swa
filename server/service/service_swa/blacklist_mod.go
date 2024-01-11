package service_swa

import (
	"github.com/ServiceWeaver/weaver"
	"time"
)

type JwtBlacklist struct {
	weaver.AutoMarshal
	ModBase
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (JwtBlacklist) TableName() string {
	return "swa_blacklist"
}

type SearchInfoBlacklist struct {
	weaver.AutoMarshal
	JwtBlacklist
	PageInfo
	StartCreatedAt time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	OrderKey       string    `json:"orderKey"`
	Desc           bool      `json:"desc"`
}
