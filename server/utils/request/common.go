package request

import "github.com/ServiceWeaver/weaver"


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
	Ids []int `json:"ids" form:"ids"`
}


type GetRoleId struct {
	RoleId uint `json:"roleId" form:"roleId"` // --角色ID
}

type Empty struct{}
