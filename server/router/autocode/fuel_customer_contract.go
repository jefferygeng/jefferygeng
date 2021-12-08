package autocode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/middleware"
)

type FuelCustomerContractRouter struct {
}

// InitFuelCustomerContractRouter 初始化 FuelCustomerContract 路由信息
func (s *FuelCustomerContractRouter) InitFuelCustomerContractRouter(Router *gin.RouterGroup) {
	fuelCustomerContractRouter := Router.Group("fuelCustomerContract").Use(middleware.OperationRecord())
	var fuelCustomerContractApi = v1.ApiGroupApp.AutoCodeApiGroup.FuelCustomerContractApi
	{
		fuelCustomerContractRouter.POST("createFuelCustomerContract", fuelCustomerContractApi.CreateFuelCustomerContract)   // 新建FuelCustomerContract
		fuelCustomerContractRouter.DELETE("deleteFuelCustomerContract", fuelCustomerContractApi.DeleteFuelCustomerContract) // 删除FuelCustomerContract
		fuelCustomerContractRouter.DELETE("deleteFuelCustomerContractByIds", fuelCustomerContractApi.DeleteFuelCustomerContractByIds) // 批量删除FuelCustomerContract
		fuelCustomerContractRouter.PUT("updateFuelCustomerContract", fuelCustomerContractApi.UpdateFuelCustomerContract)    // 更新FuelCustomerContract
		fuelCustomerContractRouter.GET("findFuelCustomerContract", fuelCustomerContractApi.FindFuelCustomerContract)        // 根据ID获取FuelCustomerContract
		fuelCustomerContractRouter.GET("getFuelCustomerContractList", fuelCustomerContractApi.GetFuelCustomerContractList)  // 获取FuelCustomerContract列表
	}
}
