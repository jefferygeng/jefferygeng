package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AirportRouter struct {
}

// InitAirportRouter 初始化 Airport 路由信息
func (s *AirportRouter) InitAirportRouter(Router *gin.RouterGroup) {
	airportRouter := Router.Group("airport").Use(middleware.OperationRecord())
	var airportApi = v1.ApiGroupApp.AutoCodeApiGroup.AirportApi
	{
		airportRouter.POST("createAirport", airportApi.CreateAirport)   // 新建Airport
		airportRouter.DELETE("deleteAirport", airportApi.DeleteAirport) // 删除Airport
		airportRouter.DELETE("deleteAirportByIds", airportApi.DeleteAirportByIds) // 批量删除Airport
		airportRouter.PUT("updateAirport", airportApi.UpdateAirport)    // 更新Airport
		airportRouter.GET("findAirport", airportApi.FindAirport)        // 根据ID获取Airport
		airportRouter.GET("getAirportList", airportApi.GetAirportList)  // 获取Airport列表
	}
}
