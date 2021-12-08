package autocode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/middleware"
)

type AgreementRouter struct {
}

// InitAgreementRouter 初始化 Agreement 路由信息
func (s *AgreementRouter) InitAgreementRouter(Router *gin.RouterGroup) {
	agreementRouter := Router.Group("agreement").Use(middleware.OperationRecord())
	var agreementApi = v1.ApiGroupApp.AutoCodeApiGroup.AgreementApi
	{
		agreementRouter.POST("createAgreement", agreementApi.CreateAgreement)   // 新建Agreement
		agreementRouter.DELETE("deleteAgreement", agreementApi.DeleteAgreement) // 删除Agreement
		agreementRouter.DELETE("deleteAgreementByIds", agreementApi.DeleteAgreementByIds) // 批量删除Agreement
		agreementRouter.PUT("updateAgreement", agreementApi.UpdateAgreement)    // 更新Agreement
		agreementRouter.GET("findAgreement", agreementApi.FindAgreement)        // 根据ID获取Agreement
		agreementRouter.GET("getAgreementList", agreementApi.GetAgreementList)  // 获取Agreement列表
	}
}
