package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/api"
	"github.com/gin-gonic/gin"
)

type OTARouter struct {
}

func (s *OTARouter) InitOTARouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.OTAApi
	{
		plugRouter.POST("build", plugApi.Build)
		plugRouter.GET("build-status", plugApi.BuildStatus)
	}
}
