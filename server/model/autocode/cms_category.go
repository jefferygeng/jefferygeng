// 自动生成模板CmsCategory
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CmsCategory 结构体
// 如果含有time.Time 请自行import time包
type CmsCategory struct {
      global.GVA_MODEL
      Categoryname  string `json:"categoryname" form:"categoryname" gorm:"column:categoryname;comment:栏目;type:varchar(20);"`
}


