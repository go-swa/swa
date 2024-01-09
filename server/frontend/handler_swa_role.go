package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func (s *server) InitRouterSwaRole(Router *gin.RouterGroup) {
	router := Router.Group("swaRole").Use(s.OperationRecord())

	router.POST("createRole", s.CreateRole)
	router.POST("deleteRole", s.DeleteRole)
	router.PUT("updateRole", s.UpdateRole)
	router.POST("setSubRoles", s.SetSubRoles)   // 设置子角色
	router.POST("setRoleMenus", s.SetRoleMenus) // 设置角色菜单

	router.POST("getRoleList", s.GetRoleList)
	router.POST("getSubRoles", s.GetSubRoles)

}


func (s *server) CreateRole(c *gin.Context) {
	var swaRole service_swa.SwaRole
	err := c.ShouldBindJSON(&swaRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(swaRole, utils.RoleVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if authBack, err := s.serviceSwa.Get().CreateRole(c, swaRole); err != nil {
		s.Logger(c).Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = s.serviceSwa.Get().AddMenuRole(c, service_swa.DefaultMenu(), swaRole.RoleId)
		_ = s.serviceSwa.Get().UpdateCasbin(c, swaRole.RoleId, service_swa.DefaultCasbin())
		response.OkWithDetailed(service_swa.SwaRoleResponse{SwaRole: authBack}, "创建成功", c)
	}
}




func (s *server) DeleteRole(c *gin.Context) {
	var swaRole service_swa.RoleUserMenu
	err := c.ShouldBindJSON(&swaRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(swaRole, utils.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteRole(c, swaRole)
	if err != nil { // --删除角色之前需要判断是否有用户正在使用此角色
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}


func (s *server) UpdateRole(c *gin.Context) {
	var auth service_swa.SwaRole
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.RoleVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	swaRole, err := s.serviceSwa.Get().UpdateRole(c, auth)
	if err != nil {
		s.Logger(c).Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(service_swa.SwaRoleResponse{SwaRole: swaRole}, "更新成功", c)
}


func (s *server) GetRoleList(c *gin.Context) {
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
	list, total, err := s.serviceSwa.Get().GetRoleList(c, pageInfo)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}


func (s *server) SetSubRoles(c *gin.Context) {
	var swaRole service_swa.SwaRole
	err := c.ShouldBindJSON(&swaRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(swaRole, utils.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().SetSubRoles(c, swaRole)
	if err != nil {
		s.Logger(c).Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

func (s *server) SetRoleMenus(c *gin.Context) {
	var auth service_swa.SwaRole
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(auth, utils.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().SetRoleMenus(c, auth)
	if err != nil {
		s.Logger(c).Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

func (s *server) GetSubRoles(c *gin.Context) {
	var roleId uint
	err := c.ShouldBindJSON(&roleId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Printf("请求子角色:\n%+v\n", roleId)

	subRoles, err := s.serviceSwa.Get().GetSubRoles(c, roleId)
	fmt.Printf("获取子角色：%v\n", subRoles)
	if err != nil {
		s.Logger(c).Error("获取子角色失败!", zap.Error(err))
		response.FailWithMessage("获取子角色失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"subRoles": subRoles}, "获取成功", c)
}
