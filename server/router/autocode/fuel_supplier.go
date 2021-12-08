package autocode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/middleware"
)

type SupplierRouter struct {
}

// InitSupplierRouter 初始化 Supplier 路由信息
func (s *SupplierRouter) InitSupplierRouter(Router *gin.RouterGroup) {
	supplierRouter := Router.Group("supplier").Use(middleware.OperationRecord())
	var supplierApi = v1.ApiGroupApp.AutoCodeApiGroup.SupplierApi
	{
		supplierRouter.POST("createSupplier", supplierApi.CreateSupplier)   // 新建Supplier
		supplierRouter.DELETE("deleteSupplier", supplierApi.DeleteSupplier) // 删除Supplier
		supplierRouter.DELETE("deleteSupplierByIds", supplierApi.DeleteSupplierByIds) // 批量删除Supplier
		supplierRouter.PUT("updateSupplier", supplierApi.UpdateSupplier)    // 更新Supplier
		supplierRouter.GET("findSupplier", supplierApi.FindSupplier)        // 根据ID获取Supplier
		supplierRouter.GET("getSupplierList", supplierApi.GetSupplierList)  // 获取Supplier列表
	}
}
