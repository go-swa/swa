package service_swa

import (
	"github.com/ServiceWeaver/weaver"
)

type ModRole struct {
	weaver.AutoMarshal
	ModBase
	RoleId        uint   `json:"roleId" gorm:"column:roleId;not null;comment:角色ID;size:90"`
	RoleName      string `json:"roleName" gorm:"column:roleName;comment:角色名"`
	ParentId      uint   `json:"parentId" gorm:"column:parentId;comment:父角色ID"`
	DefaultRouter string `json:"defaultRouter" gorm:"column:defaultRouter;comment:默认菜单;default:dashboard"`


}


type SwaRole struct {
	weaver.AutoMarshal
	ModRole

	SubRoles []ModRole `json:"subRoles" gorm:"-"`
	RoleMenus []ModMenu `json:"menus" gorm:"-"`

}

func (SwaRole) TableName() string {
	return "swa_roles"
}

type RoleUserMenu struct {
	weaver.AutoMarshal
	ModRole
	Users        []SysUser     `json:"users" gorm:"-"`
	SysBaseMenus []SysBaseMenu `json:"menus" gorm:"-"`
}

type Role2SubRoles struct {
	RoleId    uint `json:"roleId"`
	SubRoleId uint `json:"subRoleId"`
}

func (Role2SubRoles) TableName() string {
	return "swa_role_subroles"
}

type Role2Menus struct {
	RoleId uint `json:"roleId"`
	MenuId uint `json:"menuId"`
}

func (Role2Menus) TableName() string {
	return "swa_role_menus"
}

type Role2Apis struct {
	RoleId uint `json:"roleId"`
	ApiId  uint `json:"apiId"`
}

func (Role2Apis) TableName() string {
	return "swa_role_apis"
}

type Role2Users struct {
	RoleId uint `json:"roleId"`
	UserId uint `json:"userId"`
}

func (Role2Users) TableName() string {
	return "swa_role_users"
}


type SwaRoleChildren struct {
	SwaRole  `json:"swaRole"`
	Children []SwaRoleChildren `json:"children" gorm:"-"`
}

type SwaRoleResponse struct {
	SwaRole SwaRole `json:"swaRole"`
}

type SwaRoleCopyResponse struct {
	SwaRole   SwaRoleChildren `json:"swaRole"`
	OldRoleId uint            `json:"oldRoleId"` // --旧角色ID
}
