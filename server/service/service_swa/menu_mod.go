package service_swa

import (
	"github.com/ServiceWeaver/weaver"
)

type ModMenu struct {
	weaver.AutoMarshal
	ModBase
	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId" gorm:"comment:父菜单ID"`
	Path      string `json:"path" gorm:"comment:路由path"`
	Name      string `json:"name" gorm:"comment:路由name"`
	Hidden    bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort      int    `json:"sort" gorm:"comment:排序标记"`

	Meta `json:"meta" gorm:"embedded;comment:附加属性"`
}
type Meta struct {
	weaver.AutoMarshal
	ActiveName  string `json:"activeName" gorm:"comment:高亮菜单"`
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`
	CloseTab    bool   `json:"closeTab" gorm:"comment:自动关闭tab"`
}
type SysBaseMenu struct {
	weaver.AutoMarshal
	ModMenu

	SwaRoles   []SwaRole              `json:"roles" gorm:"many2many:swa_menu_role;"`
	Parameters []SysBaseMenuParameter `json:"parameters" gorm:"-"`
	MenuBtn    []SysBaseMenuBtn       `json:"menuBtn" gorm:"-"`

	Meta `json:"meta" gorm:"embedded;comment:附加属性"`
}

func (m SysBaseMenu) TableName() string {
	return "swa_menus"
}

type MenuRole struct {
	Menu       SysBaseMenu
	Roles      []SwaRole              `json:"roles" gorm:"-"`
	Parameters []SysBaseMenuParameter `json:"parameters" gorm:"-"`
	MenuBtn    []SysBaseMenuBtn       `json:"menuBtn" gorm:"-"`
}


type SysBaseMenuParameter struct {
	weaver.AutoMarshal
	ModBase
	SysBaseMenuID uint   `json:"menuId"`
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params还是query"`
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`
}

func (SysBaseMenuParameter) TableName() string {
	return "swa_menu_param"
}


type SysBaseMenuBtn struct {
	weaver.AutoMarshal
	ModBase
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}

func (SysBaseMenuBtn) TableName() string {
	return "swa_btns"
}

type SwaRoleMenu struct {
	MenuId string `json:"menuId" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	RoleId string `json:"-" gorm:"comment:角色ID;column:sys_role_role_id"`
}

func (SwaRoleMenu) TableName() string {
	return "swa_menu_role"
}

type SwaRoleBtn struct {
	RoleId           uint           `gorm:"comment:角色ID"`
	SysMenuID        uint           `gorm:"comment:菜单ID"`
	SysBaseMenuBtnID uint           `gorm:"comment:菜单按钮ID"`
	SysBaseMenuBtn   SysBaseMenuBtn ` gorm:"comment:按钮详情"`
}

func (SwaRoleBtn) TableName() string {
	return "swa_btn_role"
}

type AddMenuRoleInfo struct {
	Menus  []SysBaseMenu `json:"menus"`
	RoleId uint          `json:"roleId"`
}

func DefaultMenu() []SysBaseMenu {
	return []SysBaseMenu{{
		ModMenu: ModMenu{
			ModBase:   ModBase{ID: 1},
			ParentId:  "0",
			Path:      "dashboard",
			Name:      "dashboard",
			Component: "view/dashboard/index.vue",
			Sort:      1,
		},
		Meta: Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}

type SysMenusResponse struct {
	Menus []SysMenu `json:"menus"`
}


type SysMenu struct {
	SysBaseMenu
	MenuId     string                 `json:"menuId" gorm:"comment:菜单ID"`
	RoleId     uint                   `json:"-" gorm:"comment:角色ID"`
	Parameters []SysBaseMenuParameter `json:"parameters" gorm:"-"`
	Children   []SysMenu              `json:"children" gorm:"-"`
	Btns       map[string]uint        `json:"btns" gorm:"-"`
}

type SysBaseMenusResponse struct {
	Menus []SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu SysBaseMenu `json:"menu"`
}
