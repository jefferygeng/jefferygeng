package autocode

import (
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/autocode"
	autoCodeReq "github.com/jefferygeng/yj/server/model/autocode/request"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type FuelCustomerEmailService struct {
}

// CreateFuelCustomerEmail 创建FuelCustomerEmail记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerEmailService *FuelCustomerEmailService) CreateFuelCustomerEmail(fuelCustomerEmail autocode.FuelCustomerEmail) (err error) {
	err = global.GVA_DB.Create(&fuelCustomerEmail).Error
	return err
}

// DeleteFuelCustomerEmail 删除FuelCustomerEmail记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerEmailService *FuelCustomerEmailService)DeleteFuelCustomerEmail(fuelCustomerEmail autocode.FuelCustomerEmail) (err error) {
	err = global.GVA_DB.Delete(&fuelCustomerEmail).Error
	return err
}

// DeleteFuelCustomerEmailByIds 批量删除FuelCustomerEmail记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerEmailService *FuelCustomerEmailService)DeleteFuelCustomerEmailByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.FuelCustomerEmail{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFuelCustomerEmail 更新FuelCustomerEmail记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerEmailService *FuelCustomerEmailService)UpdateFuelCustomerEmail(fuelCustomerEmail autocode.FuelCustomerEmail) (err error) {
	err = global.GVA_DB.Save(&fuelCustomerEmail).Error
	return err
}

// GetFuelCustomerEmail 根据id获取FuelCustomerEmail记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerEmailService *FuelCustomerEmailService)GetFuelCustomerEmail(id uint) (err error, fuelCustomerEmail autocode.FuelCustomerEmail) {
	err = global.GVA_DB.Where("id = ?", id).First(&fuelCustomerEmail).Error
	return
}

// GetFuelCustomerEmailInfoList 分页获取FuelCustomerEmail记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerEmailService *FuelCustomerEmailService)GetFuelCustomerEmailInfoList(info autoCodeReq.FuelCustomerEmailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.FuelCustomerEmail{})
    var fuelCustomerEmails []autocode.FuelCustomerEmail
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.CustomerId != nil {
        db = db.Where("`customer_id` = ?",info.CustomerId)
    }
    if info.EmailType != nil {
        db = db.Where("`email_type` = ?",info.EmailType)
    }
    if info.SendType != nil {
        db = db.Where("`send_type` = ?",info.SendType)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&fuelCustomerEmails).Error
	return err, fuelCustomerEmails, total
}
