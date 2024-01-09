package service_swa

import (
	"github.com/ServiceWeaver/weaver"
	"time"
)

type ModApi struct {
	weaver.AutoMarshal
	ModBase
	AppID       uint   `json:"appID" gorm:"comment:appID"`
	TableID     uint   `json:"tableID" gorm:"comment:tableID"`
	Path        string `json:"path" gorm:"comment:api路径"`
	Description string `json:"description" gorm:"comment:api中文描述"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`
	Method      string `json:"method" gorm:"default:POST;comment:方法"`
}

func (ModApi) TableName() string {
	return "swa_apis"
}

type SearchInfoApi struct {
	weaver.AutoMarshal
	ModApi
	PageInfo
	StartCreatedAt time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	OrderKey       string    `json:"orderKey"`
	Desc           bool      `json:"desc"`
}

type SearchApiParams struct {
	ModApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc bool `json:"desc"`
}

type SysAPIResponse struct {
	Api ModApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []ModApi `json:"apis"`
}
