package frontend

import "github.com/jefferygeng/yj/server/service"

type ApiGroup struct {
	SmsApi
	SystemCfgApi
}

var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
