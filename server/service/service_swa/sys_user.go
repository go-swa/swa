package service_swa


type Register struct {
	Username  string `json:"userName" example:"用户名"`
	Password  string `json:"passWord" example:"密码"`
	NickName  string `json:"nickName" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	RoleId    uint   `json:"roleId" swaggertype:"string" example:"int 角色id"`
	Enable    int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	RoleIds   []uint `json:"roleIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone     string `json:"phone" example:"电话号码"`
	Email     string `json:"email" example:"电子邮箱"`
}


type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}


type ChangePasswordReq struct {
	ID uint `json:"-"`
	Password string `json:"password"`
	NewPassword string `json:"newPassword"`
}


type SetUserAuth struct {
	RoleId uint `json:"roleId"`
}


type SetUserAuthorities struct {
	ID      uint
	RoleIds []uint `json:"roleIds"`
}

type ChangeUserInfo struct {
	ID uint `gorm:"primarykey"`
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Phone string `json:"phone"  gorm:"comment:用户手机号"`
	RoleIds []uint `json:"roleIds" gorm:"-"`
	Email string `json:"email"  gorm:"comment:用户邮箱"`
	HeaderImg string `json:"headerImg" gorm:"comment:用户头像"`
	SideMode string `json:"sideMode"  gorm:"comment:用户侧边主题"`
	Enable int `json:"enable" gorm:"comment:冻结用户"`
	Authorities []SwaRole `json:"-" gorm:"many2many:swa_user_role;"`
}
