// 自动生成模板ZnFlight
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ZnFlight 结构体
// 如果含有time.Time 请自行import time包
type ZnFlight struct {
      global.GVA_MODEL
      Incode  string `json:"incode" form:"incode" gorm:"column:incode;comment:进港航班号;type:varchar(6);"`
      Indate  string `json:"indate" form:"indate" gorm:"column:indate;comment:进港航班日期;type:varchar(7);"`
      Iniata  string `json:"iniata" form:"iniata" gorm:"column:iniata;comment:进港航班起飞地;type:varchar(3);"`
      Intime  string `json:"intime" form:"intime" gorm:"column:intime;comment:进港航班起飞时刻;type:varchar(5);"`
      Inarrive  string `json:"inarrive" form:"inarrive" gorm:"column:inarrive;comment:进港航班落地;type:varchar(3);"`
      Inautc  string `json:"inautc" form:"inautc" gorm:"column:inautc;comment:进港航班落地时刻;type:varchar(5);"`
      Outcode  string `json:"outcode" form:"outcode" gorm:"column:outcode;comment:出港航班号;type:varchar(6);"`
      Outdate  string `json:"outdate" form:"outdate" gorm:"column:outdate;comment:出港航班日期;type:varchar(7);"`
      Outiata  string `json:"outiata" form:"outiata" gorm:"column:outiata;comment:出港航班起飞地;type:varchar(3);"`
      Outtime  string `json:"outtime" form:"outtime" gorm:"column:outtime;comment:出港航班起飞时刻;type:varchar(5);"`
      Outarrive  string `json:"outarrive" form:"outarrive" gorm:"column:outarrive;comment:出港航班落地机场;type:varchar(3);"`
      Outautc  string `json:"outautc" form:"outautc" gorm:"column:outautc;comment:出港航班落地时刻;type:varchar(5);"`
}


// TableName ZnFlight 表名
func (ZnFlight) TableName() string {
  return "zn_flight"
}

