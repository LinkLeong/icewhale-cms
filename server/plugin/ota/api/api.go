package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/service"
	"github.com/gin-gonic/gin"
)

type OTAApi struct{}

func (p *OTAApi) Build(c *gin.Context) {

	var plug = struct {
		Version     string `json:"version"`
		ReleaseNote string `json:"release_note"`
	}{}
	_ = c.ShouldBindJSON(&plug)
	go service.ServiceGroupApp.Build(plug.Version, plug.ReleaseNote)
	response.OkWithData("开始构建", c)
}

func (p *OTAApi) BuildStatus(c *gin.Context) {
	response.OkWithData(struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  service.ServiceGroupApp.Status,
		Message: service.ServiceGroupApp.Message,
	}, c)
}
