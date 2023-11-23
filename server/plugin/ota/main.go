package ota

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/router"
	"github.com/gin-gonic/gin"
)

type OTAPlugin struct {
}

func CreateOTAPlug() *OTAPlugin {
	buildHost := os.Getenv("BuildHost")
	if buildHost == "" {
		buildHost = "192.168.20.233:22"
	}
	global.GlobalConfig.BuildHost = buildHost

	global.GlobalConfig.BuildUser = "ctrdh"
	global.GlobalConfig.BuildPassword = "223306"

	buildPath := os.Getenv("BuildPath")
	if buildPath == "" {
		buildHost = "/home/ctrdh/CasaOS/operating-system"
	}
	global.GlobalConfig.BuildPath = buildPath

	return &OTAPlugin{}
}

func (*OTAPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitOTARouter(group)
}

func (*OTAPlugin) RouterPath() string {
	return "ota"
}
