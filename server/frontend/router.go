package frontend

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(ctx context.Context, s *server) *gin.Engine {
	Router := gin.Default()
	Router.StaticFS(s.swaConfig.Local.StorePath, http.Dir(s.swaConfig.Local.StorePath))

	PublicGroup := Router.Group("")
	PublicGroup.GET("/isOk?", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	s.InitRouterSwaBase(PublicGroup)

	PrivateGroup := Router.Group("")
	PrivateGroup.Use(s.JWTAuth())

	s.InitRouterConfig(PrivateGroup)
	s.InitRouterSwaRole(PrivateGroup)
	s.InitRouterSwaMenu(PrivateGroup)
	s.InitRouterSwaApi(PrivateGroup)
	s.InitRouterSwaDict(PrivateGroup)
	s.InitRouterSwaDictail(PrivateGroup)
	s.InitRouterSwaRecord(PrivateGroup)
	s.InitRouterSwaBlacklist(PrivateGroup)
	s.InitRouterSwaUser(PrivateGroup)
	s.InitRouterSwaState(PrivateGroup)
	s.InitRouterSwaCasbin(PrivateGroup)
	s.InitRouterUpDownload(PrivateGroup)

	s.Logger(ctx).Debug("router register success")
	return Router
}
