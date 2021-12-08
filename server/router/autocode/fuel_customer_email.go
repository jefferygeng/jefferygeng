package autocode

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jefferygeng/yj/server/api/v1"
	"github.com/jefferygeng/yj/server/middleware"
)

type FuelCustomerEmailRouter struct {
}

// InitFuelCustomerEmailRouter 初始化 FuelCustomerEmail 路由信息
func (s *FuelCustomerEmailRouter) InitFuelCustomerEmailRouter(Router *gin.RouterGroup) {
	fuelCustomerEmailRouter := Router.Group("fuelCustomerEmail").Use(middleware.OperationRecord())
	var fuelCustomerEmailApi = v1.ApiGroupApp.AutoCodeApiGroup.FuelCustomerEmailApi
	{
		fuelCustomerEmailRouter.POST("createFuelCustomerEmail", fuelCustomerEmailApi.CreateFuelCustomerEmail)             // 新建FuelCustomerEmail
		fuelCustomerEmailRouter.DELETE("deleteFuelCustomerEmail", fuelCustomerEmailApi.DeleteFuelCustomerEmail)           // 删除FuelCustomerEmail
		fuelCustomerEmailRouter.DELETE("deleteFuelCustomerEmailByIds", fuelCustomerEmailApi.DeleteFuelCustomerEmailByIds) // 批量删除FuelCustomerEmail
		fuelCustomerEmailRouter.PUT("updateFuelCustomerEmail", fuelCustomerEmailApi.UpdateFuelCustomerEmail)              // 更新FuelCustomerEmail
		fuelCustomerEmailRouter.GET("findFuelCustomerEmail", fuelCustomerEmailApi.FindFuelCustomerEmail)                  // 根据ID获取FuelCustomerEmail
		fuelCustomerEmailRouter.GET("getFuelCustomerEmailList", fuelCustomerEmailApi.GetFuelCustomerEmailList)            // 获取FuelCustomerEmail列表
	}
}
