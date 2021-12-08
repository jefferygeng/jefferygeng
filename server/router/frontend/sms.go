package frontend

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type SmsRouter struct {
}

// InitSmsRouter 初始化 Airport 路由信息
func (s *SmsRouter) InitSmsRouter(Router *gin.RouterGroup) {
	smsRouter := Router.Group("sms")
	var smsApi = v1.ApiGroupApp.FrontendApiGroup.SmsApi
	{
		smsRouter.GET("sendCode", smsApi.SendMobileSms) // 发送验证短信
		// smsRouter.DELETE("deleteAirport", airportApi.DeleteAirport)           // 删除Airport
		// smsRouter.DELETE("deleteAirportByIds", airportApi.DeleteAirportByIds) // 批量删除Airport
		// smsRouter.PUT("updateAirport", airportApi.UpdateAirport)              // 更新Airport
		// smsRouter.GET("findAirport", airportApi.FindAirport)                  // 根据ID获取Airport
		// smsRouter.GET("getAirportList", airportApi.GetAirportList)            // 获取Airport列表
	}
}
