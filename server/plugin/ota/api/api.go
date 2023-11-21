package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type OTAApi struct{}

func (p *OTAApi) Build(c *gin.Context) {
	response.OkWithData("开始构建", c)
}
