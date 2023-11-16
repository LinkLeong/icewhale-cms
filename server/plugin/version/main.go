package version

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/router"
	"github.com/gin-gonic/gin"
)

type VersionPlugin struct {
}

func CreateVersionPlug() *VersionPlugin {
	global.GlobalConfig.Url = "http://origin.casaos.io/casaos-api"
	return &VersionPlugin{}
}

func (*VersionPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitVersionRouter(group)
}

func (*VersionPlugin) RouterPath() string {
	return "version"
}
