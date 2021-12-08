package autocode

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jefferygeng/yj/server/api/v1"
	"github.com/jefferygeng/yj/server/middleware"
)

type CmsContentRouter struct {
}

// InitCmsContentRouter 初始化 CmsContent 路由信息
func (s *CmsContentRouter) InitCmsContentRouter(Router *gin.RouterGroup) {
	cmsContentRouter := Router.Group("cmsContent").Use(middleware.OperationRecord())
	var cmsContentApi = v1.ApiGroupApp.AutoCodeApiGroup.CmsContentApi
	{
		cmsContentRouter.POST("createCmsContent", cmsContentApi.CreateCmsContent)             // 新建CmsContent
		cmsContentRouter.DELETE("deleteCmsContent", cmsContentApi.DeleteCmsContent)           // 删除CmsContent
		cmsContentRouter.DELETE("deleteCmsContentByIds", cmsContentApi.DeleteCmsContentByIds) // 批量删除CmsContent
		cmsContentRouter.PUT("updateCmsContent", cmsContentApi.UpdateCmsContent)              // 更新CmsContent
		cmsContentRouter.GET("findCmsContent", cmsContentApi.FindCmsContent)                  // 根据ID获取CmsContent
		cmsContentRouter.GET("getCmsContentList", cmsContentApi.GetCmsContentList)            // 获取CmsContent列表
	}
}
