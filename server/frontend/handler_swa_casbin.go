package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"strings"
)


func (s *server) InitRouterSwaCasbin(Router *gin.RouterGroup) {
	router := Router.Group("casbin").Use(s.OperationRecord())

	router.POST("updateCasbin", s.UpdateCasbin)
	router.POST("getCasbinByRoleId", s.GetCasbinByRoleId)
}


func (s *server) UpdateCasbin(c *gin.Context) {
	var cmr service_swa.CasbinInReceive
	err := c.ShouldBindJSON(&cmr)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(cmr, utils.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = s.serviceSwa.Get().UpdateCasbin(c, cmr.RoleId, cmr.CasbinInfos)
	if err != nil {
		s.Logger(c).Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}


func (s *server) GetCasbinByRoleId(c *gin.Context) {
	var casbin service_swa.CasbinInReceive
	err := c.ShouldBindJSON(&casbin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(casbin, utils.RoleIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	paths, _ := s.serviceSwa.Get().GetPolicyPathByRoleId(c, casbin.RoleId)
	response.OkWithDetailed(service_swa.PolicyPathResponse{Paths: paths}, "获取成功", c)
}

func (s *server) CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if s.swaConfig.System.Env != "develop" {
			waitUse, _ := utils.GetClaims(c, s.swaConfig.Jwt.SigningKey)
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, s.swaConfig.System.RouterPrefix)
			act := c.Request.Method
			sub := strconv.Itoa(int(waitUse.RoleId))
			sok, err := s.serviceSwa.Get().Scasbin(c, sub, obj, act)
			if err != nil {
				fmt.Printf("获取casbin结果错误，err:%v", err)
			}
			if !sok {
				response.FailWithDetailed(gin.H{}, "权限不足casbin", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
