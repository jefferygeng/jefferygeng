package autocode

import (
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/autocode"
	autoCodeReq "github.com/jefferygeng/yj/server/model/autocode/request"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type FuelCustomerContactService struct {
}

// CreateFuelCustomerContact 创建FuelCustomerContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContactService *FuelCustomerContactService) CreateFuelCustomerContact(fuelCustomerContact autocode.FuelCustomerContact) (err error) {
	err = global.GVA_DB.Create(&fuelCustomerContact).Error
	return err
}

// DeleteFuelCustomerContact 删除FuelCustomerContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContactService *FuelCustomerContactService)DeleteFuelCustomerContact(fuelCustomerContact autocode.FuelCustomerContact) (err error) {
	err = global.GVA_DB.Delete(&fuelCustomerContact).Error
	return err
}

// DeleteFuelCustomerContactByIds 批量删除FuelCustomerContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContactService *FuelCustomerContactService)DeleteFuelCustomerContactByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.FuelCustomerContact{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFuelCustomerContact 更新FuelCustomerContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContactService *FuelCustomerContactService)UpdateFuelCustomerContact(fuelCustomerContact autocode.FuelCustomerContact) (err error) {
	err = global.GVA_DB.Save(&fuelCustomerContact).Error
	return err
}

// GetFuelCustomerContact 根据id获取FuelCustomerContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContactService *FuelCustomerContactService)GetFuelCustomerContact(id uint) (err error, fuelCustomerContact autocode.FuelCustomerContact) {
	err = global.GVA_DB.Where("id = ?", id).First(&fuelCustomerContact).Error
	return
}

// GetFuelCustomerContactInfoList 分页获取FuelCustomerContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContactService *FuelCustomerContactService)GetFuelCustomerContactInfoList(info autoCodeReq.FuelCustomerContactSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.FuelCustomerContact{})
    var fuelCustomerContacts []autocode.FuelCustomerContact
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fuelCustomerContacts).Error
	return err, fuelCustomerContacts, total
}
