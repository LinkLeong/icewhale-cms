package api

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/version/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VersionApi struct{}

// @Tags Version
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /version/routerName [post]
func (p *VersionApi) GetList(c *gin.Context) {

	var plug model.Request
	_ = c.ShouldBindJSON(&plug)

	if res, err := service.ServiceGroupApp.GetVersion(); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(res, c)
	}

}
func (p *VersionApi) Delete(c *gin.Context) {
	id := c.Param("id")
	if res, err := service.ServiceGroupApp.DeleteVersion(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(fmt.Sprintf("删除失败！%v", err), c)
	} else {
		response.OkWithData(res, c)
	}
}
func (p *VersionApi) AddVersion(c *gin.Context) {

	var plug model.Version
	_ = c.ShouldBindJSON(&plug)

	if res, err := service.ServiceGroupApp.AddVersion(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(res, c)
	}

}
