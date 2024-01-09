package service_file

import (
	"github.com/ServiceWeaver/weaver"
	"time"
)

type ModBase struct {
	weaver.AutoMarshal
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index" json:"-"`
	CreatedBy uint      `json:"createdBy" gorm:"column:createdBy;comment:创建者"`
	UpdatedBy uint      `json:"updatedBy" gorm:"column:updatedBy;comment:更新者"`
	DeletedBy uint      `json:"deletedBy" gorm:"column:deletedBy"`
}

type PageInfo struct {
	weaver.AutoMarshal
	Page     int    `json:"page" form:"page"`         // --页码
	PageSize int    `json:"pageSize" form:"pageSize"` // --每页大小
	Keyword  string `json:"keyword" form:"keyword"`   // --关键字
}


type GetById struct {
	ID int `json:"id" form:"id"` // --主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []uint `json:"ids" form:"ids"`
}


type GetRoleId struct {
	RoleId uint `json:"roleId" form:"roleId"` // --角色ID
}

type Empty struct{}
