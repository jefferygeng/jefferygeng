// 自动生成模板Airport
package autocode

import (
	"github.com/jefferygeng/yj/server/global"
)

// Airport 结构体
// 如果含有time.Time 请自行import time包
type Airport struct {
	global.GVA_MODEL
	Iata       string   `json:"iata" form:"iata" gorm:"column:iata;comment:IATA;type:varchar(10);"`
	Icao       string   `json:"icao" form:"icao" gorm:"column:icao;comment:ICAO;type:varchar(10);"`
	Cname      string   `json:"cname" form:"cname" gorm:"column:cname;comment:机场中文名称;type:varchar(50);"`
	Ename      string   `json:"ename" form:"ename" gorm:"column:ename;comment:机场英文名称;type:varchar(100);"`
	Country    string   `json:"country" form:"country" gorm:"column:country;comment:国家;type:varchar(50);"`
	Supplierid uint     `json:"supplierid" form:"supplierid" gorm:"column:supplierid;comment:供应商;type:bigint"`
	Region     *int     `json:"region" form:"region" gorm:"column:region;comment:所属地区"`
	Source     *int     `json:"source" form:"source" gorm:"column:source;comment:机场归属;type:int"`
	Supplier   Supplier `gorm:"foreignkey:Supplierid"`
}

// TableName Airport 表名
func (Airport) TableName() string {
	return "fuel_airport"
}
