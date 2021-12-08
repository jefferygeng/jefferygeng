// 自动生成模板Supplier
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Supplier 结构体
// 如果含有time.Time 请自行import time包
type Supplier struct {
      global.GVA_MODEL
      Supplier_code  string `json:"supplier_code" form:"supplier_code" gorm:"column:supplier_code;comment:供应商编号;type:varchar(16);"`
      Supplier_cn  string `json:"supplier_cn" form:"supplier_cn" gorm:"column:supplier_cn;comment:公司中文名称;type:varchar(50);"`
      Supplier_en  string `json:"supplier_en" form:"supplier_en" gorm:"column:supplier_en;comment:公司英文名称;type:varchar(50);"`
      Supplier_small  string `json:"supplier_small" form:"supplier_small" gorm:"column:supplier_small;comment:公司名称缩写;type:varchar(15);"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:公司地址;type:varchar(100);"`
      Source  *int `json:"source" form:"source" gorm:"column:source;comment:是否境内供应商;type:smallint"`
}


// TableName Supplier 表名
func (Supplier) TableName() string {
  return "fuel_supplier"
}

