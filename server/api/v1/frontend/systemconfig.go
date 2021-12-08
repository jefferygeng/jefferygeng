package frontend

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemCfgApi struct {
}

//var SendDAO = service.ServiceGroupApp.FrontendServiceGroup.SmsCodeService

// @Tags app接口
// @Summary 获取app公用配置信息
// @Produce  application/json
// @Success 200 {string} json "{"code":0,"data":{"Logo":"/logo.jpg","Mobile":"+86 (10) 6409 4980","Email":"rsea@royal-sinoenergy.com","Appname":"中能天盛","Servicecontent":"暂无","Version":"v1.0"},"msg":"获取成功"}}"
// @Router /systemcfg/configInfo [GET]
func (scapi *SystemCfgApi) SystemConfigInfo(c *gin.Context) {
	if err, config := systemConfigService.GetSystemConfig(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		result := systemRes.SysConfigResponse{Config: config}
		response.OkWithDetailed(result.Config.AppInfo, "获取成功", c)
	}

}
