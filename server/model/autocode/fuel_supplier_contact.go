// 自动生成模板FuelSupplierContact
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FuelSupplierContact 结构体
// 如果含有time.Time 请自行import time包
type FuelSupplierContact struct {
      global.GVA_MODEL
      Truename  string `json:"truename" form:"truename" gorm:"column:truename;comment:姓名;type:varchar(16);"`
      Job  string `json:"job" form:"job" gorm:"column:job;comment:职位/负责事务;type:varchar(50);"`
      Mobile  string `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机;type:varchar(30);"`
      Tel  string `json:"tel" form:"tel" gorm:"column:tel;comment:座机;type:varchar(30);"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:邮箱;type:varchar(60);"`
      BindId  *int `json:"bind_id" form:"bind_id" gorm:"column:bind_id;comment:绑定的供应商ID;type:bigint"`
      Supplier Supplier  `gorm:"foreignkey:BindId"`

}


// TableName FuelSupplierContact 表名
func (FuelSupplierContact) TableName() string {
  return "fuel_supplier_contact"
}

//type Fuelsupplier struct{
//      global.GVA_MODEL
//      Id uint
//      Supplier_code  string `json:"supplier_code" form:"supplier_code" gorm:"column:supplier_code;comment:供应商编号;type:varchar(16);"`
//      Supplier_cn  string `json:"supplier_cn" form:"supplier_cn" gorm:"column:supplier_cn;comment:公司中文名称;type:varchar(50);"`
//      Supplier_en  string `json:"supplier_en" form:"supplier_en" gorm:"column:supplier_en;comment:公司英文名称;type:varchar(50);"`
//      Supplier_small  string `json:"supplier_small" form:"supplier_small" gorm:"column:supplier_small;comment:公司名称缩写;type:varchar(15);"`
//}