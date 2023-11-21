package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ota/service"
	"github.com/gin-gonic/gin"
)

type OTAApi struct{}

func (p *OTAApi) Build(c *gin.Context) {
	go service.ServiceGroupApp.Build()
	response.OkWithData("开始构建", c)
}
