package frontend

import (
	"demo1/service/service_config"
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *server) InitRouterSwaState(Router *gin.RouterGroup) {
	router := Router.Group("system").Use(s.OperationRecord())
	routerNoRecord := Router.Group("system")

	router.POST("getSystemConfig", s.GetSystemConfig)
	router.POST("setSystemConfig", s.SetSystemConfig)
	router.POST("reloadSystem", s.ReloadSystem)

	routerNoRecord.POST("getServerInfo", s.GetServerInfo)
}


func (s *server) GetSystemConfig(c *gin.Context) {
	config, err := s.serviceConfig.Get().GetCfgSwaConfig(c)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"config": config}, "获取成功", c)
}


func (s *server) SetSystemConfig(c *gin.Context) {
	var sys service_config.ModSwaConfig
	err := c.ShouldBindJSON(&sys)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceConfig.Get().SetSystemConfig(c, sys)
	if err != nil {
		s.Logger(c).Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}


func (s *server) ReloadSystem(c *gin.Context) {
	err := utils.Reload()
	if err != nil {
		s.Logger(c).Error("重启系统失败!", zap.Error(err))
		response.FailWithMessage("重启系统失败", c)
		return
	}
	response.OkWithMessage("重启系统成功", c)
}


func (s *server) GetServerInfo(c *gin.Context) {
	serverInfo, err := s.serviceSwa.Get().GetServerInfo(c)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"server": serverInfo, "stateDict": service_swa.StateDict}, "获取成功", c)
}
