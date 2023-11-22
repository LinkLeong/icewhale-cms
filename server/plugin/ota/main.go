package ota

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/router"
	"github.com/gin-gonic/gin"
)

type OTAPlugin struct {
}

func CreateOTAPlug() *OTAPlugin {
	global.GlobalConfig.BuildHost = "192.168.20.233:22"
	global.GlobalConfig.BuildUser = "ctrdh"
	global.GlobalConfig.BuildPassword = "223306"
	global.GlobalConfig.BuildPath = "/home/ctrdh/operating-system"
	return &OTAPlugin{}
}

func (*OTAPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitOTARouter(group)
}

func (*OTAPlugin) RouterPath() string {
	return "ota"
}
