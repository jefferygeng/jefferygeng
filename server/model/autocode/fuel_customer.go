// 自动生成模板FuelCustomer
package autocode

import (
	"github.com/jefferygeng/yj/server/global"
)

// FuelCustomer 结构体
// 如果含有time.Time 请自行import time包
type FuelCustomer struct {
	global.GVA_MODEL
	CustomerCode      string `json:"customerCode" form:"customerCode" gorm:"column:customer_code;comment:客户编号;type:varchar(20);"`
	Cname             string `json:"cname" form:"cname" gorm:"column:cname;comment:公司中文全称;type:varchar(50);"`
	Ename             string `json:"ename" form:"ename" gorm:"column:ename;comment:公司英文全称;type:varchar(50);"`
	ShortName         string `json:"shortName" form:"shortName" gorm:"column:short_name;comment:公司名称缩写;type:varchar(20);"`
	Address           string `json:"address" form:"address" gorm:"column:address;comment:公司地址;type:varchar(150);"`
	RegAddress        string `json:"regAddress" form:"regAddress" gorm:"column:reg_address;comment:公司注册地;type:varchar(150);"`
	RegisteredCapital string `json:"registeredCapital" form:"registeredCapital" gorm:"column:registered_capital;comment:注册资金;type:varchar(20);"`
	Iata              string `json:"iata" form:"iata" gorm:"column:iata;comment:二字码;type:varchar(2);"`
	Icao              string `json:"icao" form:"icao" gorm:"column:icao;comment:三字码;type:varchar(3);"`
	Ctype             *int   `json:"ctype" form:"ctype" gorm:"column:ctype;comment:客户性质;type:int"`
}

// TableName FuelCustomer 表名
func (FuelCustomer) TableName() string {
	return "fuel_customer"
}
