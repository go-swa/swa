package service_swa

import (
	"github.com/ServiceWeaver/weaver"
	"time"
)

type ModDictail struct {
	weaver.AutoMarshal
	ModBase
	Label     string `json:"label" form:"label" gorm:"column:label;comment:展示值"`
	Value     int    `json:"value" form:"value" gorm:"column:value;comment:字典值"`
	Status    *bool  `json:"status" form:"status" gorm:"column:status;comment:启用状态"`
	Sort      int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序标记"`
	SwaDictID int    `json:"swaDictID" form:"swaDictID" gorm:"column:swaDictID;comment:关联标记"`
}

func (ModDictail) TableName() string {
	return "swa_dict_details"
}

type SearchInfoDictail struct {
	weaver.AutoMarshal
	ModDictail
	PageInfo
	StartCreatedAt time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	OrderKey       string    `json:"orderKey"`
	Desc           bool      `json:"desc"`
}

type SwaDictailSearch struct {
	weaver.AutoMarshal
	ModDictail
	PageInfo
}
