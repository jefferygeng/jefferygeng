package autocode

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jefferygeng/yj/server/api/v1"
	"github.com/jefferygeng/yj/server/middleware"
)

type FuelCustomerContactRouter struct {
}

// InitFuelCustomerContactRouter 初始化 FuelCustomerContact 路由信息
func (s *FuelCustomerContactRouter) InitFuelCustomerContactRouter(Router *gin.RouterGroup) {
	fuelCustomerContactRouter := Router.Group("fuelCustomerContact").Use(middleware.OperationRecord())
	var fuelCustomerContactApi = v1.ApiGroupApp.AutoCodeApiGroup.FuelCustomerContactApi
	{
		fuelCustomerContactRouter.POST("createFuelCustomerContact", fuelCustomerContactApi.CreateFuelCustomerContact)             // 新建FuelCustomerContact
		fuelCustomerContactRouter.DELETE("deleteFuelCustomerContact", fuelCustomerContactApi.DeleteFuelCustomerContact)           // 删除FuelCustomerContact
		fuelCustomerContactRouter.DELETE("deleteFuelCustomerContactByIds", fuelCustomerContactApi.DeleteFuelCustomerContactByIds) // 批量删除FuelCustomerContact
		fuelCustomerContactRouter.PUT("updateFuelCustomerContact", fuelCustomerContactApi.UpdateFuelCustomerContact)              // 更新FuelCustomerContact
		fuelCustomerContactRouter.GET("findFuelCustomerContact", fuelCustomerContactApi.FindFuelCustomerContact)                  // 根据ID获取FuelCustomerContact
		fuelCustomerContactRouter.GET("getFuelCustomerContactList", fuelCustomerContactApi.GetFuelCustomerContactList)            // 获取FuelCustomerContact列表
	}
}
