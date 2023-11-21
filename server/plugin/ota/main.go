package version

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/router"
	"github.com/gin-gonic/gin"
)

type OTAPlugin struct {
}

func CreateVersionPlug() *OTAPlugin {
	global.GlobalConfig.Url = "http://origin.casaos.io/casaos-api"
	return &OTAPlugin{}
}

func (*OTAPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitVersionRouter(group)
}

func (*OTAPlugin) RouterPath() string {
	return "release"
}
