// 自动生成模板FuelCustomerEmail
package autocode

import (
	"github.com/jefferygeng/yj/server/global"
)

// FuelCustomerEmail 结构体
// 如果含有time.Time 请自行import time包
type FuelCustomerEmail struct {
	global.GVA_MODEL
	CustomerId   *int   `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:绑定的客户的ID;type:bigint"`
	EmailType    *int   `json:"emailType" form:"emailType" gorm:"column:email_type;comment:0为价格邮件 1 为账单邮件;type:int"`
	SendType     *int   `json:"sendType" form:"sendType" gorm:"column:send_type;comment:发送类型 0为发送 1 为抄送;type:int"`
	EmailContent string `json:"emailContent" form:"emailContent" gorm:"column:email_content;comment:email地址;type:varchar(50);"`
}

// TableName FuelCustomerEmail 表名
func (FuelCustomerEmail) TableName() string {
	return "fuel_customer_email"
}
