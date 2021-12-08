package autocode

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jefferygeng/yj/server/api/v1"
	"github.com/jefferygeng/yj/server/middleware"
)

type AirportRouter struct {
}

// InitAirportRouter 初始化 Airport 路由信息
func (s *AirportRouter) InitAirportRouter(Router *gin.RouterGroup) {
	airportRouter := Router.Group("airport").Use(middleware.OperationRecord())
	var airportApi = v1.ApiGroupApp.AutoCodeApiGroup.AirportApi
	{
		airportRouter.POST("createAirport", airportApi.CreateAirport)             // 新建Airport
		airportRouter.DELETE("deleteAirport", airportApi.DeleteAirport)           // 删除Airport
		airportRouter.DELETE("deleteAirportByIds", airportApi.DeleteAirportByIds) // 批量删除Airport
		airportRouter.PUT("updateAirport", airportApi.UpdateAirport)              // 更新Airport
		airportRouter.GET("findAirport", airportApi.FindAirport)                  // 根据ID获取Airport
		airportRouter.GET("getAirportList", airportApi.GetAirportList)            // 获取Airport列表
	}
}
