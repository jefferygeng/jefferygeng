package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FuelCustomerContactRouter struct {
}

// InitFuelCustomerContactRouter 初始化 FuelCustomerContact 路由信息
func (s *FuelCustomerContactRouter) InitFuelCustomerContactRouter(Router *gin.RouterGroup) {
	fuelCustomerContactRouter := Router.Group("fuelCustomerContact").Use(middleware.OperationRecord())
	var fuelCustomerContactApi = v1.ApiGroupApp.AutoCodeApiGroup.FuelCustomerContactApi
	{
		fuelCustomerContactRouter.POST("createFuelCustomerContact", fuelCustomerContactApi.CreateFuelCustomerContact)   // 新建FuelCustomerContact
		fuelCustomerContactRouter.DELETE("deleteFuelCustomerContact", fuelCustomerContactApi.DeleteFuelCustomerContact) // 删除FuelCustomerContact
		fuelCustomerContactRouter.DELETE("deleteFuelCustomerContactByIds", fuelCustomerContactApi.DeleteFuelCustomerContactByIds) // 批量删除FuelCustomerContact
		fuelCustomerContactRouter.PUT("updateFuelCustomerContact", fuelCustomerContactApi.UpdateFuelCustomerContact)    // 更新FuelCustomerContact
		fuelCustomerContactRouter.GET("findFuelCustomerContact", fuelCustomerContactApi.FindFuelCustomerContact)        // 根据ID获取FuelCustomerContact
		fuelCustomerContactRouter.GET("getFuelCustomerContactList", fuelCustomerContactApi.GetFuelCustomerContactList)  // 获取FuelCustomerContact列表
	}
}
