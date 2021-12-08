// 自动生成模板FuelCustomerContract
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FuelCustomerContract 结构体
// 如果含有time.Time 请自行import time包
type FuelCustomerContract struct {
      global.GVA_MODEL
      ContractCode  string `json:"contractCode" form:"contractCode" gorm:"column:contract_code;comment:合同编号;type:varchar(50);"`
      ContractExpire  string `json:"contractExpire" form:"contractExpire" gorm:"column:contract_expire;comment:合同有效期;type:varchar(50);"`
      ContractFilepath  string `json:"contractFilepath" form:"contractFilepath" gorm:"column:contract_filepath;comment:合同附件;type:varchar(50);"`
      Bindgs  *int `json:"bindgs" form:"bindgs" gorm:"column:bindgs;comment:与之签订合同的公司;type:varchar(50);"`
      Currency  *int `json:"currency" form:"currency" gorm:"column:currency;comment:币种;type:int"`
      Customer_id  *int `json:"customer_id" form:"customer_id" gorm:"column:customer_id;comment:绑定的客户id;type:int"`
      FuelCustomer FuelCustomer `gorm:"foreignkey:Customer_id"`
}


// TableName FuelCustomerContract 表名
func (FuelCustomerContract) TableName() string {
  return "fuel_customer_contract"
}

