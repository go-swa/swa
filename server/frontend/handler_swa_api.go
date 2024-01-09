package frontend

import (
	"demo1/service/service_swa"
	"demo1/utils"
	"demo1/utils/request"
	"demo1/utils/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func (s *server) InitRouterSwaApi(Router *gin.RouterGroup) {
	router := Router.Group("api").Use(s.OperationRecord())
	routerNoRecord := Router.Group("api")

	router.POST("createApi", s.CreateApi)
	router.POST("deleteApi", s.DeleteApi)
	router.POST("getApiById", s.GetApiById)
	router.POST("updateApi", s.UpdateApi)
	router.DELETE("deleteApisByIds", s.DeleteApisByIds)

	routerNoRecord.POST("getAllApis", s.GetAllApis)
	routerNoRecord.POST("getApiList", s.GetApiList)
}


func (s *server) CreateApi(c *gin.Context) {
	var api service_swa.ModApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().CreateApi(c, api)
	if err != nil {
		s.Logger(c).Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}


func (s *server) DeleteApi(c *gin.Context) {
	var api service_swa.ModApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(service_swa.ModBase{}, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteApi(c, api)
	if err != nil {
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}


func (s *server) GetApiList(c *gin.Context) {
	var pageInfo service_swa.SearchApiParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := s.serviceSwa.Get().GetAPIInfoList(c, pageInfo.ModApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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


func (s *server) GetApiById(c *gin.Context) {
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
	api, err := s.serviceSwa.Get().GetApiById(c, idInfo.ID)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(service_swa.SysAPIResponse{Api: api}, "获取成功", c)
}


func (s *server) UpdateApi(c *gin.Context) {
	var api service_swa.ModApi
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().UpdateApi(c, api)
	if err != nil {
		s.Logger(c).Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}


func (s *server) GetAllApis(c *gin.Context) {
	apis, err := s.serviceSwa.Get().GetAllApis(c)
	if err != nil {
		s.Logger(c).Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(service_swa.SysAPIListResponse{Apis: apis}, "获取成功", c)
}


func (s *server) DeleteApisByIds(c *gin.Context) {
	var ids service_swa.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = s.serviceSwa.Get().DeleteApisByIds(c, ids.Ids)
	if err != nil {
		s.Logger(c).Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
