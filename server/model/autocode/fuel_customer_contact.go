// 自动生成模板FuelCustomerContact
package autocode

import (
	"github.com/jefferygeng/yj/server/global"
)

// FuelCustomerContact 结构体
// 如果含有time.Time 请自行import time包
type FuelCustomerContact struct {
	global.GVA_MODEL
	Truename string `json:"truename" form:"truename" gorm:"column:truename;comment:姓名;type:varchar(16);"`
	Job      string `json:"job" form:"job" gorm:"column:job;comment:职位/负责事务;type:varchar(50);"`
	Mobile   string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机;type:varchar(30);"`
	Tel      string `json:"tel" form:"tel" gorm:"column:tel;comment:座机;type:varchar(30);"`
	Email    string `json:"email" form:"email" gorm:"column:email;comment:邮箱;type:varchar(60);"`
	BindId   *int   `json:"bindId" form:"bindId" gorm:"column:bind_id;comment:绑定的客户ID;type:bigint"`
}

// TableName FuelCustomerContact 表名
func (FuelCustomerContact) TableName() string {
	return "fuel_customer_contact"
}
