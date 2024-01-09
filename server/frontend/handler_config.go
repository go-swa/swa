package frontend

import (
	"demo1/service/service_config"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
)


func (s *server) InitRouterConfig(Router *gin.RouterGroup) {
	routerNoRecord := Router.Group("config")
	router := Router.Group("config").Use(s.OperationRecord())

	router.POST("setConfigGormDb", s.SetConfigGormDb)

	routerNoRecord.GET("getConfigSwaConfig", s.GetConfigSwaConfig)
	routerNoRecord.GET("getConfigGormDb", s.GetConfigGormDb)
	routerNoRecord.GET("getConfigJwt", s.GetConfigJwt)
}


func (s *server) GetConfigSwaConfig(c *gin.Context) {
	ms, err := s.serviceConfig.Get().GetCfgSwaConfig(c)
	if err != nil {
		response.FailWithMessage("初始化失败", c)
	} else {
		response.OkWithDetailed(ms, "初始化成功", c)
	}
}

func (s *server) SetConfigGormDb(c *gin.Context) {
	var m service_config.ModGormDb
	var err error

	err = c.ShouldBindJSON(&m)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = s.serviceConfig.Get().SetCfgGormDb(c, m)
	if err != nil {
		response.FailWithMessage("create fail", c)
	} else {
		response.OkWithDetailed(m, "create success", c)
	}
}

func (s *server) GetConfigGormDb(c *gin.Context) {
	ms, err := s.serviceConfig.Get().GetCfgGormDb(c)
	if err != nil {
		response.FailWithMessage("get gorm db config fail", c)
	} else {
		response.OkWithDetailed(ms, "成功", c)
	}
}

func (s *server) GetConfigJwt(c *gin.Context) {
	ms, err := s.serviceConfig.Get().GetCfgJwt(c)
	if err != nil {
		response.FailWithMessage("get jwt config fail", c)
	} else {
		response.OkWithDetailed(ms, "成功", c)
	}
}
