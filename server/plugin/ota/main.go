package ota

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/router"
	"github.com/gin-gonic/gin"
)

type OTAPlugin struct {
}

func CreateOTAPlug() *OTAPlugin {
	global.GlobalConfig.BuildHost = "192.168.20.233:22"
	global.GlobalConfig.BuildUser = "ctrdh"
	global.GlobalConfig.BuildPassword = "223306"
	return &OTAPlugin{}
}

func (*OTAPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitVersionRouter(group)
}

func (*OTAPlugin) RouterPath() string {
	return "ota"
}
