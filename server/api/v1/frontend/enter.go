package frontend

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SmsApi
	SystemCfgApi
}

var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
