package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func (s *server) InitRouterSwaDict(Router *gin.RouterGroup) {
	routerRecord := Router.Group("swaDict")
	router := Router.Group("swaDict").Use(s.OperationRecord())

	router.POST("createSwaDict", s.CreateSwaDict)
	router.DELETE("deleteSwaDict", s.DeleteSwaDict)
	router.PUT("updateSwaDict", s.UpdateSwaDict)

	routerRecord.GET("findSwaDict", s.FindSwaDict)
	routerRecord.GET("getSwaDictList", s.GetSwaDictList)
}


func (s *server) CreateSwaDict(c *gin.Context) {
	var dictionary service_swa.ModDict
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().CreateSwaDict(c, dictionary)
	if err != nil {
		s.Logger(c).Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}


func (s *server) DeleteSwaDict(c *gin.Context) {
	var dictionary service_swa.ModDict
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteSwaDict(c, dictionary)
	if err != nil {
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}


func (s *server) UpdateSwaDict(c *gin.Context) {
	var dictionary service_swa.ModDict
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().UpdateSwaDict(c, dictionary)
	if err != nil {
		s.Logger(c).Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}


func (s *server) FindSwaDict(c *gin.Context) {
	var dictionary service_swa.ModDict
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	swaDict, err := s.serviceSwa.Get().GetSwaDict(c, dictionary.Type, dictionary.ID, dictionary.Status)
	if err != nil {
		s.Logger(c).Error("字典未创建或未开启!", zap.Error(err))
		response.FailWithMessage("字典未创建或未开启", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysDictionary": swaDict}, "查询成功", c)
}


func (s *server) GetSwaDictList(c *gin.Context) {
	var pageInfo service_swa.SwaDictSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := s.serviceSwa.Get().GetSwaDictInfoList(c, pageInfo)
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
