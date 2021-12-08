package utils

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/jefferygeng/yj/server/global"
	"go.uber.org/zap"
)

type Code struct {
	viCode string `json:"code"`
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendSms(mobile string, code string) (_err error) {

	//判断手机号是否已经发过验证码
	// smsCfg := &config.Sms{
	// 	AccessKeyId:     global.GVA_CONFIG.Sms.AccessKeyId,
	// 	AccessKeySecret: global.GVA_CONFIG.Sms.AccessKeySecret,
	// 	SignName:        global.GVA_CONFIG.Sms.SignName,
	// 	TemplateCode:    global.GVA_CONFIG.Sms.TemplateCode,
	// }
	client, _err := CreateClient(tea.String(""+global.GVA_CONFIG.Sms.AccessKeyId+""), tea.String(global.GVA_CONFIG.Sms.AccessKeySecret))
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &mobile,
		SignName:      tea.String(global.GVA_CONFIG.Sms.SignName),
		TemplateCode:  tea.String(global.GVA_CONFIG.Sms.TemplateCode),
		TemplateParam: tea.String("{\"code\" : " + code + "}"),
	}
	resp, _err := client.SendSms(sendSmsRequest)
	if _err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", _err))
		return _err
	}
	respJson := util.ToJSONString(tea.ToMap(resp.Body))
	global.GVA_LOG.Info("SUCCESS", zap.Any("info", respJson))
	codeStatus := resp.Body.Code
	if *util.ToString([]byte(*codeStatus)) == "OK" {
		return nil
	}
	return nil
}
