package frontend

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type SystemCfgRouter struct {
}

// InitSystemCfgRouter 初始化app所需要的用户信息
func (s *SystemCfgRouter) InitSystemCfgRouter(Router *gin.RouterGroup) {
	scRouter := Router.Group("systemcfg")
	var scApi = v1.ApiGroupApp.FrontendApiGroup.SystemCfgApi
	{
		scRouter.GET("configInfo", scApi.SystemConfigInfo) // 获取系统配置信息
	}
}
