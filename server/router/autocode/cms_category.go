package autocode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/middleware"
)

type CmsCategoryRouter struct {
}

// InitCmsCategoryRouter 初始化 CmsCategory 路由信息
func (s *CmsCategoryRouter) InitCmsCategoryRouter(Router *gin.RouterGroup) {
	cmsCategoryRouter := Router.Group("cmsCategory").Use(middleware.OperationRecord())
	var cmsCategoryApi = v1.ApiGroupApp.AutoCodeApiGroup.CmsCategoryApi
	{
		cmsCategoryRouter.POST("createCmsCategory", cmsCategoryApi.CreateCmsCategory)   // 新建CmsCategory
		cmsCategoryRouter.DELETE("deleteCmsCategory", cmsCategoryApi.DeleteCmsCategory) // 删除CmsCategory
		cmsCategoryRouter.DELETE("deleteCmsCategoryByIds", cmsCategoryApi.DeleteCmsCategoryByIds) // 批量删除CmsCategory
		cmsCategoryRouter.PUT("updateCmsCategory", cmsCategoryApi.UpdateCmsCategory)    // 更新CmsCategory
		cmsCategoryRouter.GET("findCmsCategory", cmsCategoryApi.FindCmsCategory)        // 根据ID获取CmsCategory
		cmsCategoryRouter.GET("getCmsCategoryList", cmsCategoryApi.GetCmsCategoryList)  // 获取CmsCategory列表
	}
}
