package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/api"
	"github.com/gin-gonic/gin"
)

type VersionRouter struct {
}

func (s *VersionRouter) InitVersionRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.VersionApi
	{
		//plugRouter.POST("", plugApi.ApiName)
		plugRouter.GET("list", plugApi.GetList)
		plugRouter.DELETE(":id", plugApi.Delete)
		plugRouter.POST("add", plugApi.AddVersion)
	}
}
