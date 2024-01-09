package service_swa

import "github.com/ServiceWeaver/weaver"

type CasbinInfo struct {
	weaver.AutoMarshal
	Path   string `json:"path"`
	Method string `json:"method"`
}


type CasbinInReceive struct {
	RoleId      uint         `json:"roleId"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/menu/getMenu", Method: "POST"},
		{Path: "/jwt/jsonInBlacklist", Method: "POST"},
		{Path: "/base/login", Method: "POST"},
		{Path: "/user/admin_register", Method: "POST"},
		{Path: "/user/changePassword", Method: "POST"},
		{Path: "/user/setUserRole", Method: "POST"},
		{Path: "/user/setUserInfo", Method: "PUT"},
		{Path: "/user/getUserInfo", Method: "GET"},
	}
}

type PolicyPathResponse struct {
	Paths []CasbinInfo `json:"paths"`
}
