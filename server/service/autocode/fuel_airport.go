package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AirportService struct {
}

// CreateAirport 创建Airport记录
// Author [piexlmax](https://github.com/piexlmax)
func (airportService *AirportService) CreateAirport(airport autocode.Airport) (err error) {
	err = global.GVA_DB.Create(&airport).Error
	return err
}

// DeleteAirport 删除Airport记录
// Author [piexlmax](https://github.com/piexlmax)
func (airportService *AirportService)DeleteAirport(airport autocode.Airport) (err error) {
	err = global.GVA_DB.Delete(&airport).Error
	return err
}

// DeleteAirportByIds 批量删除Airport记录
// Author [piexlmax](https://github.com/piexlmax)
func (airportService *AirportService)DeleteAirportByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Airport{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAirport 更新Airport记录
// Author [piexlmax](https://github.com/piexlmax)
func (airportService *AirportService)UpdateAirport(airport autocode.Airport) (err error) {
	err = global.GVA_DB.Save(&airport).Error
	return err
}

// GetAirport 根据id获取Airport记录
// Author [piexlmax](https://github.com/piexlmax)
func (airportService *AirportService)GetAirport(id uint) (err error, airport autocode.Airport) {
	err = global.GVA_DB.Where("id = ?", id).First(&airport).Error
	return
}

// GetAirportInfoList 分页获取Airport记录
// Author [piexlmax](https://github.com/piexlmax)
func (airportService *AirportService)GetAirportInfoList(info autoCodeReq.AirportSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Airport{})
    var airports []autocode.Airport
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Iata != "" {
        db = db.Where("`iata` = ?",info.Iata)
    }
    if info.Icao != "" {
        db = db.Where("`icao` = ?",info.Icao)
    }
    if info.Cname != "" {
        db = db.Where("`cname` LIKE ?","%"+ info.Cname+"%")
    }
    if info.Ename != "" {
        db = db.Where("`ename` LIKE ?","%"+ info.Ename+"%")
    }
    if info.Country != "" {
        db = db.Where("`country` = ?",info.Country)
    }
    if info.Region != nil {
        db = db.Where("`region` = ?",info.Region)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Preload("Supplier").Offset(offset).Find(&airports).Error
	return err, airports, total
}
