package frontend

import (
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/common/response"
	"github.com/jefferygeng/yj/server/service"
	"github.com/jefferygeng/yj/server/utils"
	"go.uber.org/zap"
)

type SmsApi struct {
}

var SendDAO = service.ServiceGroupApp.FrontendServiceGroup.SmsCodeService

// @Tags app接口
// @Summary 发送手机验证码
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"发送成功"}"
// @Router /sms/sendCode [GET]
func (smsapi *SmsApi) SendMobileSms(c *gin.Context) {
	codeName, isbool := c.GetQuery("code")

	if !isbool {
		response.FailWithMessage("缺少code参数", c)
		return
	}
	if codeName == "" {
		response.FailWithMessage("code参数不能为空", c)
		return
	}
	verifycode := utils.GenValidateCode(6)
	if err := utils.SendSms(codeName, verifycode); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
		return
	} else {

		_err := SendDAO.SendCodeDAO(codeName, verifycode)
		if _err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Any("err", _err))
			response.OkWithMessage("发送失败", c)
		}
		response.OkWithMessage("发送成功", c)
	}

}
