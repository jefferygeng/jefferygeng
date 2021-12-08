package autocode

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jefferygeng/yj/server/api/v1"
	"github.com/jefferygeng/yj/server/middleware"
)

type FuelCustomerRouter struct {
}

// InitFuelCustomerRouter 初始化 FuelCustomer 路由信息
func (s *FuelCustomerRouter) InitFuelCustomerRouter(Router *gin.RouterGroup) {
	fuelCustomerRouter := Router.Group("fuelCustomer").Use(middleware.OperationRecord())
	var fuelCustomerApi = v1.ApiGroupApp.AutoCodeApiGroup.FuelCustomerApi
	{
		fuelCustomerRouter.POST("createFuelCustomer", fuelCustomerApi.CreateFuelCustomer)             // 新建FuelCustomer
		fuelCustomerRouter.DELETE("deleteFuelCustomer", fuelCustomerApi.DeleteFuelCustomer)           // 删除FuelCustomer
		fuelCustomerRouter.DELETE("deleteFuelCustomerByIds", fuelCustomerApi.DeleteFuelCustomerByIds) // 批量删除FuelCustomer
		fuelCustomerRouter.PUT("updateFuelCustomer", fuelCustomerApi.UpdateFuelCustomer)              // 更新FuelCustomer
		fuelCustomerRouter.GET("findFuelCustomer", fuelCustomerApi.FindFuelCustomer)                  // 根据ID获取FuelCustomer
		fuelCustomerRouter.GET("getFuelCustomerList", fuelCustomerApi.GetFuelCustomerList)            // 获取FuelCustomer列表
	}
}
