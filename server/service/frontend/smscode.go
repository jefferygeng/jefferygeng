package frontend

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/frontend"
)

type SmsCodeService struct {
}

//@author: yj
//@function: SendCodeDAO
//@description: 发送验证码后入库保存记录
//@param:
//@return: error

func (e *SmsCodeService) SendCodeDAO(phone string, code string) (err error) {
	var smModel frontend.SmsCode
	smModel.Phone = phone
	smModel.Code = code
	err = global.GVA_DB.Create(&smModel).Error
	return err
}

  func (e *SmsCodeService) ValidateSmsCode(phone string, code string) (err error,smsDetail *frontend.SmsCode ){
    var smModel frontend.SmsCode
	if err := global.GVA_DB.Where("phone = ? and code = ?",phone,code).First(&smModel).Error; err != nil{
			return errors.New("没有找到此验证码"),nil
	}
	  return nil,&smModel
  }