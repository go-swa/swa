package service_swa

type SwaRoleBtnReq struct {
	MenuID   uint   `json:"menuID"`
	RoleId   uint   `json:"roleId"`
	Selected []uint `json:"selected"`
}

type SwaRoleBtnRes struct {
	Selected []uint `json:"selected"`
}
