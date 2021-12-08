package frontend

import "github.com/flipped-aurora/gin-vue-admin/server/global"

//验证码表结构
type SmsCode struct {
	global.GVA_MODEL
	Phone string `gorm:"varchar(11)" json:"phone"`
	Code  string `gorm:"varchar(4)"  json:"code"`
}

// TableName SmsCode 表名
func (SmsCode) TableName() string {
	return "zn_smscode"
}
