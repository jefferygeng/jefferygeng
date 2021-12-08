package autocode

import (
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/autocode"
	autoCodeReq "github.com/jefferygeng/yj/server/model/autocode/request"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type FuelSupplierContactService struct {
}

// CreateFuelSupplierContact 创建FuelSupplierContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelSupplierContactService *FuelSupplierContactService) CreateFuelSupplierContact(fuelSupplierContact autocode.FuelSupplierContact) (err error) {
	err = global.GVA_DB.Create(&fuelSupplierContact).Error
	return err
}

// DeleteFuelSupplierContact 删除FuelSupplierContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelSupplierContactService *FuelSupplierContactService) DeleteFuelSupplierContact(fuelSupplierContact autocode.FuelSupplierContact) (err error) {
	err = global.GVA_DB.Delete(&fuelSupplierContact).Error
	return err
}

// DeleteFuelSupplierContactByIds 批量删除FuelSupplierContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelSupplierContactService *FuelSupplierContactService) DeleteFuelSupplierContactByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.FuelSupplierContact{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFuelSupplierContact 更新FuelSupplierContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelSupplierContactService *FuelSupplierContactService) UpdateFuelSupplierContact(fuelSupplierContact autocode.FuelSupplierContact) (err error) {
	err = global.GVA_DB.Save(&fuelSupplierContact).Error
	return err
}

// GetFuelSupplierContact 根据id获取FuelSupplierContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelSupplierContactService *FuelSupplierContactService) GetFuelSupplierContact(id uint) (err error, fuelSupplierContact autocode.FuelSupplierContact) {
	err = global.GVA_DB.Where("id = ?", id).First(&fuelSupplierContact).Error
	return
}

// GetFuelSupplierContactInfoList 分页获取FuelSupplierContact记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelSupplierContactService *FuelSupplierContactService) GetFuelSupplierContactInfoList(info autoCodeReq.FuelSupplierContactSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.FuelSupplierContact{})
	var fuelSupplierContacts []autocode.FuelSupplierContact
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Supplier").Find(&fuelSupplierContacts).Error
	return err, fuelSupplierContacts, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSupplierName
//@description: 返回当前选中suppliername
//@param: id float64

func (fuelSupplierContactService *FuelSupplierContactService) GetSupplierName(id float64) (err error, fuelSupplierContact autocode.FuelSupplierContact) {
	err = global.GVA_DB.Preload("Suppliers").Where("id = ?", id).First(&fuelSupplierContact).Error
	return
}
