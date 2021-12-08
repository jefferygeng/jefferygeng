// 自动生成模板CmsContent
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CmsContent 结构体
// 如果含有time.Time 请自行import time包
type CmsContent struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;type:varchar(50);"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`
      Aliasname  string `json:"aliasname" form:"aliasname" gorm:"column:aliasname;comment:标题别名;type:varchar(20);"`
}


// TableName CmsContent 表名
func (CmsContent) TableName() string {
  return "cmsContent"
}

