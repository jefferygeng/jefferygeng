package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type FuelCustomerService struct {
}

// CreateFuelCustomer 创建FuelCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerService *FuelCustomerService) CreateFuelCustomer(fuelCustomer autocode.FuelCustomer) (err error) {
	err = global.GVA_DB.Create(&fuelCustomer).Error
	return err
}

// DeleteFuelCustomer 删除FuelCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerService *FuelCustomerService)DeleteFuelCustomer(fuelCustomer autocode.FuelCustomer) (err error) {
	err = global.GVA_DB.Delete(&fuelCustomer).Error
	return err
}

// DeleteFuelCustomerByIds 批量删除FuelCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerService *FuelCustomerService)DeleteFuelCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.FuelCustomer{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFuelCustomer 更新FuelCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerService *FuelCustomerService)UpdateFuelCustomer(fuelCustomer autocode.FuelCustomer) (err error) {
	err = global.GVA_DB.Save(&fuelCustomer).Error
	return err
}

// GetFuelCustomer 根据id获取FuelCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerService *FuelCustomerService)GetFuelCustomer(id uint) (err error, fuelCustomer autocode.FuelCustomer) {
	err = global.GVA_DB.Where("id = ?", id).First(&fuelCustomer).Error
	return
}

// GetFuelCustomerInfoList 分页获取FuelCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerService *FuelCustomerService)GetFuelCustomerInfoList(info autoCodeReq.FuelCustomerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.FuelCustomer{})
    var fuelCustomers []autocode.FuelCustomer
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.CustomerCode != "" {
        db = db.Where("`customer_code` = ?",info.CustomerCode)
    }
    if info.Cname != "" {
        db = db.Where("`cname` LIKE ?","%"+ info.Cname+"%")
    }
    if info.Ename != "" {
        db = db.Where("`ename` LIKE ?","%"+ info.Ename+"%")
    }
    if info.ShortName != "" {
        db = db.Where("`short_name` = ?",info.ShortName)
    }
    if info.Address != "" {
        db = db.Where("`address` LIKE ?","%"+ info.Address+"%")
    }
    if info.Iata != "" {
        db = db.Where("`iata` = ?",info.Iata)
    }
    if info.Icao != "" {
        db = db.Where("`icao` = ?",info.Icao)
    }
    if info.Ctype != nil {
        db = db.Where("`ctype` = ?",info.Ctype)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fuelCustomers).Error
	return err, fuelCustomers, total
}
