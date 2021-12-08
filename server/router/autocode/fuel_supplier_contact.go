package autocode

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jefferygeng/yj/server/api/v1"
	"github.com/jefferygeng/yj/server/middleware"
)

type FuelSupplierContactRouter struct {
}

// InitFuelSupplierContactRouter 初始化 FuelSupplierContact 路由信息
func (s *FuelSupplierContactRouter) InitFuelSupplierContactRouter(Router *gin.RouterGroup) {
	fuelSupplierContactRouter := Router.Group("fuelSupplierContact").Use(middleware.OperationRecord())
	var fuelSupplierContactApi = v1.ApiGroupApp.AutoCodeApiGroup.FuelSupplierContactApi
	{
		fuelSupplierContactRouter.POST("createFuelSupplierContact", fuelSupplierContactApi.CreateFuelSupplierContact)             // 新建FuelSupplierContact
		fuelSupplierContactRouter.DELETE("deleteFuelSupplierContact", fuelSupplierContactApi.DeleteFuelSupplierContact)           // 删除FuelSupplierContact
		fuelSupplierContactRouter.DELETE("deleteFuelSupplierContactByIds", fuelSupplierContactApi.DeleteFuelSupplierContactByIds) // 批量删除FuelSupplierContact
		fuelSupplierContactRouter.PUT("updateFuelSupplierContact", fuelSupplierContactApi.UpdateFuelSupplierContact)              // 更新FuelSupplierContact
		fuelSupplierContactRouter.GET("findFuelSupplierContact", fuelSupplierContactApi.FindFuelSupplierContact)                  // 根据ID获取FuelSupplierContact
		fuelSupplierContactRouter.GET("getFuelSupplierContactList", fuelSupplierContactApi.GetFuelSupplierContactList)            // 获取FuelSupplierContact列表
	}
}
