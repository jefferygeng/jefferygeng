package autocode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/middleware"
)

type ZnFlightRouter struct {
}

// InitZnFlightRouter 初始化 ZnFlight 路由信息
func (s *ZnFlightRouter) InitZnFlightRouter(Router *gin.RouterGroup) {
	znFlightRouter := Router.Group("znFlight").Use(middleware.OperationRecord())
	var znFlightApi = v1.ApiGroupApp.AutoCodeApiGroup.ZnFlightApi
	{
		znFlightRouter.POST("createZnFlight", znFlightApi.CreateZnFlight)   // 新建ZnFlight
		znFlightRouter.DELETE("deleteZnFlight", znFlightApi.DeleteZnFlight) // 删除ZnFlight
		znFlightRouter.DELETE("deleteZnFlightByIds", znFlightApi.DeleteZnFlightByIds) // 批量删除ZnFlight
		znFlightRouter.PUT("updateZnFlight", znFlightApi.UpdateZnFlight)    // 更新ZnFlight
		znFlightRouter.GET("findZnFlight", znFlightApi.FindZnFlight)        // 根据ID获取ZnFlight
		znFlightRouter.GET("getZnFlightList", znFlightApi.GetZnFlightList)  // 获取ZnFlight列表
	}
}
