package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func (s *server) InitRouterSwaDictail(Router *gin.RouterGroup) {
	routerNoRecord := Router.Group("swaDictail")
	router := Router.Group("swaDictail").Use(s.OperationRecord())

	router.POST("createSwaDictail", s.CreateSwaDictail)
	router.DELETE("deleteSwaDictail", s.DeleteSwaDictail)
	router.PUT("updateSwaDictail", s.UpdateSwaDictail)

	routerNoRecord.GET("findSwaDictail", s.FindSwaDictail)
	routerNoRecord.GET("getSwaDictailList", s.GetSwaDictailList)
}


func (s *server) CreateSwaDictail(c *gin.Context) {
	var detail service_swa.ModDictail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().CreateSwaDictail(c, detail)
	if err != nil {
		s.Logger(c).Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}


func (s *server) DeleteSwaDictail(c *gin.Context) {
	var detail service_swa.ModDictail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteSwaDictail(c, detail)
	if err != nil {
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}


func (s *server) UpdateSwaDictail(c *gin.Context) {
	var detail service_swa.ModDictail
	err := c.ShouldBindJSON(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().UpdateSwaDictail(c, detail)
	if err != nil {
		s.Logger(c).Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}


func (s *server) FindSwaDictail(c *gin.Context) {
	var detail service_swa.ModDictail
	err := c.ShouldBindQuery(&detail)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(detail, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSwaDictail, err := s.serviceSwa.Get().GetSwaDictail(c, detail.ID)
	if err != nil {
		s.Logger(c).Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSwaDictail": reSwaDictail}, "查询成功", c)
}


func (s *server) GetSwaDictailList(c *gin.Context) {
	var pageInfo service_swa.SwaDictailSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := s.serviceSwa.Get().GetSwaDictailInfoList(c, pageInfo)
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
