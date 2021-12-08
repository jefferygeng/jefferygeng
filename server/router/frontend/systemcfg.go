package frontend

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jefferygeng/yj/server/api/v1"
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
