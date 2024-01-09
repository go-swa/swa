package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/request"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)


func (s *server) InitRouterSwaUser(Router *gin.RouterGroup) {
	routerNoRecord := Router.Group("user")
	router := Router.Group("user").Use(s.OperationRecord())

	router.POST("admin_register", s.Register)               // --管理员注册账号
	router.POST("changePassword", s.ChangePassword)         // --用户修改密码
	router.POST("setUserRole", s.SetUserRole)               // --设置用户权限
	router.DELETE("deleteUser", s.DeleteUser)               // --删除用户
	router.PUT("setUserInfo", s.SetUserInfo)                // --设置用户信息
	router.PUT("setSelfInfo", s.SetSelfInfo)                // --设置自身信息
	router.POST("setUserAuthorities", s.SetUserAuthorities) // --设置用户权限组
	router.POST("resetPassword", s.ResetPassword)           // --设置用户权限组
	router.GET("getUserInfo", s.GetUserInfo)                // --获取自身信息
	router.GET("getUserDict", s.GetUserDict)                // --获取用户字典

	routerNoRecord.POST("getUserList", s.GetUserList) // --分页获取用户列表

}


type SysUserResponse struct {
	User service_swa.SysUser `json:"user"`
}

type ChangePasswordReq struct {
	ID uint `json:"-"`
	Password string `json:"password"`
	NewPassword string `json:"newPassword"`
}


func (s *server) Register(c *gin.Context) {
	var m service_swa.SysUser
	err := c.ShouldBindJSON(&m)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(m, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userReturn, err := s.serviceSwa.Get().Register(c, m)
	if err != nil {
		s.Logger(c).Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(SysUserResponse{User: userReturn}, "注册成功", c)
}


func (s *server) ChangePassword(c *gin.Context) {
	var req ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c, s.swaConfig.Jwt.SigningKey)
	_, err = s.serviceSwa.Get().ChangePassword(c, uid, req.Password, req.NewPassword)
	if err != nil {
		s.Logger(c).Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}


func (s *server) GetUserList(c *gin.Context) {
	var pageInfo service_swa.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := s.serviceSwa.Get().GetUserInfoList(c, pageInfo)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}


func (s *server) SetUserRole(c *gin.Context) {
	var sua service_swa.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserRoleVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c, s.swaConfig.Jwt.SigningKey)
	err = s.serviceSwa.Get().SetUserRole(c, userID, sua.RoleId)
	if err != nil {
		s.Logger(c).Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c, s.swaConfig.Jwt.SigningKey)
	j := utils.NewJWT(s.swaConfig.Jwt.SigningKey)
	claims.RoleId = sua.RoleId
	if token, err := j.CreateToken(*claims); err != nil {
		s.Logger(c).Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
		response.OkWithMessage("修改成功", c)
	}
}


func (s *server) SetUserAuthorities(c *gin.Context) {
	var sua service_swa.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().SetUserAuthorities(c, sua.ID, sua.RoleIds)
	if err != nil {
		s.Logger(c).Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}


func (s *server) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	jwtId := s.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 自杀失败", c)
		return
	}
	err = s.serviceSwa.Get().DeleteUser(c, reqId.ID)
	if err != nil {
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}


func (s *server) SetUserInfo(c *gin.Context) {
	var user service_swa.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.RoleIds) != 0 {
		err = s.serviceSwa.Get().SetUserAuthorities(c, user.ID, user.RoleIds)
		if err != nil {
			s.Logger(c).Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
			return
		}
	}
	err = s.serviceSwa.Get().SetUserInfo(c, service_swa.SysUser{
		ModBase: service_swa.ModBase{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		s.Logger(c).Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}


func (s *server) SetSelfInfo(c *gin.Context) {
	var user service_swa.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = s.GetUserID(c)
	err = s.serviceSwa.Get().SetSelfInfo(c, service_swa.SysUser{
		ModBase: service_swa.ModBase{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		s.Logger(c).Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}



func (s *server) ResetPassword(c *gin.Context) {
	var user service_swa.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().ResetPassword(c, user.ID)
	if err != nil {
		s.Logger(c).Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}


func (s *server) GetUserDict(c *gin.Context) {
	userDict, err := s.serviceSwa.Get().GetUserDict(c)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"userDict": userDict}, "获取成功", c)
	}
}


func (s *server) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c, s.swaConfig.Jwt.SigningKey)

	user, err := s.serviceSwa.Get().GetUserInfo(c, uuid)
	if err != nil {
		s.Logger(c).Error("s 获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("s 获取用户信息失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": user}, "获取用户信息成功", c)
}

func (s *server) GetUserID(c *gin.Context) uint {
	id := utils.GetUserID(c, s.swaConfig.Jwt.SigningKey)
	return id
}
