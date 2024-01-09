package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/request"
	"demo1/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)


func (s *server) InitRouterSwaMenu(Router *gin.RouterGroup) {
	routerNoRecord := Router.Group("menu")
	router := Router.Group("menu").Use(s.OperationRecord())

	router.POST("addBaseMenu", s.AddBaseMenu)
	router.POST("addMenuRole", s.AddMenuRole)
	router.POST("deleteBaseMenu", s.DeleteBaseMenu)
	router.POST("updateBaseMenu", s.UpdateBaseMenu)

	routerNoRecord.POST("getMenu", s.GetMenu)
	routerNoRecord.POST("getMenuList", s.GetMenuList)
	routerNoRecord.POST("getBaseMenuTree", s.GetBaseMenuTree)
	routerNoRecord.POST("getMenuRole", s.GetMenuRole)
	routerNoRecord.POST("getBaseMenuById", s.GetBaseMenuById)
}

func (s *server) GetMenu(c *gin.Context) {
	var allMenus []service_swa.SysMenu
	var menus []service_swa.SysMenu
	roleId := s.GetUserRoleId(c)
	if roleId < 1 {
		s.Logger(c).Error("获取角色ID小于1")
		response.FailWithMessage("获取失败, 角色id为0", c)
		return
	}

	roleMenus, err := s.serviceSwa.Get().GetRoleMenu(c, roleId)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	for i := range roleMenus {
		allMenus = append(allMenus, service_swa.SysMenu{
			SysBaseMenu: roleMenus[i],
			RoleId:      roleId,
			MenuId:      strconv.Itoa(int(roleMenus[i].ID)),
			Parameters:  roleMenus[i].Parameters,
		})
	}

	menuTree := make(map[string][]service_swa.SysMenu)
	for _, v := range allMenus {
		menuTree[v.ParentId] = append(menuTree[v.ParentId], v)
	}
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = s.getChildrenList(&menus[i], menuTree)
	}
	response.OkWithDetailed(service_swa.SysMenusResponse{Menus: menus}, "获取成功", c)
}

func (s *server) getChildrenList(menu *service_swa.SysMenu, treeMap map[string][]service_swa.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for j := 0; j < len(menu.Children); j++ {
		err = s.getChildrenList(&menu.Children[j], treeMap)
	}
	return err
}


func (s *server) GetUserRoleId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := utils.GetClaims(c, s.swaConfig.Jwt.SigningKey); err != nil {
			return 0
		} else {
			return cl.RoleId
		}
	} else {
		waitUse := claims.(*utils.CustomClaims)
		return waitUse.RoleId
	}
}


type TreeMenu struct {
	service_swa.SysBaseMenu
	Children []TreeMenu `json:"children" gorm:"-"`
}

func (s *server) GetMenuList(c *gin.Context) {
	var treeMenu []TreeMenu
	var pageInfo request.PageInfo
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
	roleMenus, total, err := s.serviceSwa.Get().GetInfoList(c)
	if err != nil {
		s.Logger(c).Error("获取菜单失败!", zap.Error(err))
		response.FailWithMessage("获取菜单失败", c)
		return
	}
	for i := range roleMenus {
		treeMenu = append(treeMenu, TreeMenu{
			SysBaseMenu: roleMenus[i],
		})
	}

	treeMap := make(map[string][]TreeMenu)
	for _, v := range treeMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	var menuList []TreeMenu
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = s.getBaseChildrenList(&menuList[i], treeMap)
	}

	response.OkWithDetailed(response.PageResult{
		List:     menuList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}


func (s *server) getBaseChildrenList(menu *TreeMenu, treeMap map[string][]TreeMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = s.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}


func (s *server) GetBaseMenuTree(c *gin.Context) {
	var treeMenu []TreeMenu
	allMenus, err := s.serviceSwa.Get().GetAllBaseMenu(c)
	if err != nil {
		s.Logger(c).Error("获取菜单失败!", zap.Error(err))
		response.FailWithMessage("获取菜单失败", c)
		return
	}

	for i := range allMenus {
		treeMenu = append(treeMenu, TreeMenu{
			SysBaseMenu: allMenus[i],
		})
	}

	treeMap := make(map[string][]TreeMenu)
	for _, v := range treeMenu {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}

	var menus []TreeMenu
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = s.getBaseChildrenList(&menus[i], treeMap)
	}

	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
}


func (s *server) AddMenuRole(c *gin.Context) {
	var roleMenu service_swa.AddMenuRoleInfo
	err := c.ShouldBindJSON(&roleMenu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(roleMenu, utils.RoleIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := s.serviceSwa.Get().AddMenuRole(c, roleMenu.Menus, roleMenu.RoleId); err != nil {
		s.Logger(c).Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}


func (s *server) GetMenuRole(c *gin.Context) {
	var param request.GetRoleId
	var menus []service_swa.SysMenu
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(param, utils.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	baseMenu, err := s.serviceSwa.Get().GetMenuRole(c, param.RoleId)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithDetailed(service_swa.SysMenusResponse{Menus: menus}, "获取失败", c)
		return
	}

	for i := range baseMenu {
		menus = append(menus, service_swa.SysMenu{
			SysBaseMenu: baseMenu[i],
			RoleId:      param.RoleId,
			MenuId:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}

	response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
}


func (s *server) AddBaseMenu(c *gin.Context) {
	var menu service_swa.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().AddBaseMenu(c, menu)
	if err != nil {
		s.Logger(c).Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}


func (s *server) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetById
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uuid := utils.GetUserUuid(c, s.swaConfig.Jwt.SigningKey)
	user, err := s.serviceSwa.Get().GetUserInfo(c, uuid)
	if err != nil {
		s.Logger(c).Error("删除失败:", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("删除失败:%v", err), c)
		return
	}
	if user.Username != "superadmin" {
		response.FailWithMessage(fmt.Sprintf("删除失败:只有超级管理员才有权限删除菜单"), c)
		return
	}
	err = utils.Verify(menu, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteBaseMenu(c, menu.ID)
	if err != nil {
		s.Logger(c).Error("删除失败:", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("删除失败:%v", err), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}


func (s *server) UpdateBaseMenu(c *gin.Context) {
	var menu service_swa.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu.Meta, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().UpdateBaseMenu(c, menu)
	if err != nil {
		s.Logger(c).Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}


func (s *server) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menu, err := s.serviceSwa.Get().GetBaseMenuById(c, idInfo.ID)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(service_swa.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
}

