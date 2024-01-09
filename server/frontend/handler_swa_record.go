package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func (s *server) InitRouterSwaRecord(Router *gin.RouterGroup) {
	router := Router.Group("swaRecord").Use(s.OperationRecord())
	routerNoRecord := Router.Group("swaRecord")

	router.POST("createSwaRecord", s.CreateSwaRecord)
	router.DELETE("deleteSwaRecord", s.DeleteSwaRecord)
	router.DELETE("deleteSwaRecordByIds", s.DeleteSwaRecordByIds)

	routerNoRecord.GET("findSwaRecord", s.FindSwaRecord)
	routerNoRecord.GET("getSwaRecordList", s.GetSwaRecordList)
}


func (s *server) CreateSwaRecord(c *gin.Context) {
	var swaRecord service_swa.SwaRecord
	err := c.ShouldBindJSON(&swaRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().CreateSwaRecord(c, swaRecord)
	if err != nil {
		s.Logger(c).Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}


func (s *server) DeleteSwaRecord(c *gin.Context) {
	var swaRecord service_swa.SwaRecord
	err := c.ShouldBindJSON(&swaRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteSwaRecord(c, swaRecord)
	if err != nil {
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}


func (s *server) DeleteSwaRecordByIds(c *gin.Context) {
	var IDS service_swa.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteSwaRecordByIds(c, IDS.Ids)
	if err != nil {
		s.Logger(c).Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}


func (s *server) FindSwaRecord(c *gin.Context) {
	var swaRecord service_swa.SwaRecord
	err := c.ShouldBindQuery(&swaRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(swaRecord, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	reSwaRecord, err := s.serviceSwa.Get().GetSwaRecord(c, swaRecord.ID)
	if err != nil {
		s.Logger(c).Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"reSwaRecord": reSwaRecord}, "查询成功", c)
}


func (s *server) GetSwaRecordList(c *gin.Context) {
	var pageInfo service_swa.SwaRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := s.serviceSwa.Get().GetSwaRecordInfoList(c, pageInfo)
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
