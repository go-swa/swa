package service_swa

import (
	"github.com/ServiceWeaver/weaver"
	"github.com/gofrs/uuid"
)

type SysUser struct {
	weaver.AutoMarshal
	ModBase
	UUID        uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	Username    string    `json:"userName" gorm:"index;comment:用户登录名"`
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`
	HeaderImg   string    `json:"headerImg" gorm:"comment:用户头像"`
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`
	Phone       string    `json:"phone"  gorm:"comment:用户手机号"`
	Email       string    `json:"email"  gorm:"comment:用户邮箱"`
	Team        string    `json:"team"  gorm:"comment:项目组"`
	Dept        string    `json:"dept"  gorm:"comment:部门"`
	Company     string    `json:"company"  gorm:"comment:公司/组织"`
	Enable      int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`

	RoleId      uint      `json:"roleId" gorm:"default:888;comment:用户角色ID"`
	RoleIds     []uint    `json:"roleIds" gorm:"-"`
	SwaRole     SwaRole   `json:"swaRole" gorm:"-"`
	Authorities []SwaRole `json:"authorities" gorm:"-"`
}

func (SysUser) TableName() string {
	return "swa_users"
}






type UserRole struct {
	UserId uint `gorm:"column:user_id"`
	RoleId uint `gorm:"column:role_id"`
}

func (s *UserRole) TableName() string {
	return "swa_user_role"
}

type SysUserResponse struct {
	User SysUser `json:"user"`
}

type LoginResponse struct {
	User      SysUser `json:"user"`
	Token     string  `json:"token"`
	ExpiresAt int64   `json:"expiresAt"`
}
