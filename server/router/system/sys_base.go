package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("applogin", baseApi.AppLogin) //app登录专用
		baseRouter.POST("smslogin", baseApi.SmsLogin)
		baseRouter.POST("captcha", baseApi.Captcha)
	}

}
