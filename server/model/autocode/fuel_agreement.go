// 自动生成模板Agreement
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Agreement 结构体
// 如果含有time.Time 请自行import time包
type Agreement struct {
      global.GVA_MODEL
      Sncode  string `json:"sncode" form:"sncode" gorm:"column:sncode;comment:合同编号;type:varchar(50);"`
      Dateout  string `json:"dateout" form:"dateout" gorm:"column:dateout;comment:合同有效期;type:varchar(100);"`
      Filespath  string `json:"filespath" form:"filespath" gorm:"column:filespath;comment:合同附件;type:varchar(255);"`
      Bindgs  *int `json:"bindgs" form:"bindgs" gorm:"column:bindgs;comment:签订公司;type:int"`
      Currency  *int `json:"currency" form:"currency" gorm:"column:currency;comment:币种;type:int"`
      Bindsupplier  *int `json:"bindsupplier" form:"bindsupplier" gorm:"column:bindsupplier;comment:绑定的代理商;type:bigint"`
      Supplier Supplier  `gorm:"foreignkey:Bindsupplier"`
}


// TableName Agreement 表名
func (Agreement) TableName() string {
  return "fuel_agreement"
}

