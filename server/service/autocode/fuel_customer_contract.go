package autocode

import (
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/autocode"
	autoCodeReq "github.com/jefferygeng/yj/server/model/autocode/request"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type FuelCustomerContractService struct {
}

// CreateFuelCustomerContract 创建FuelCustomerContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContractService *FuelCustomerContractService) CreateFuelCustomerContract(fuelCustomerContract autocode.FuelCustomerContract) (err error) {
	err = global.GVA_DB.Create(&fuelCustomerContract).Error
	return err
}

// DeleteFuelCustomerContract 删除FuelCustomerContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContractService *FuelCustomerContractService) DeleteFuelCustomerContract(fuelCustomerContract autocode.FuelCustomerContract) (err error) {
	err = global.GVA_DB.Delete(&fuelCustomerContract).Error
	return err
}

// DeleteFuelCustomerContractByIds 批量删除FuelCustomerContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContractService *FuelCustomerContractService) DeleteFuelCustomerContractByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.FuelCustomerContract{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFuelCustomerContract 更新FuelCustomerContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContractService *FuelCustomerContractService) UpdateFuelCustomerContract(fuelCustomerContract autocode.FuelCustomerContract) (err error) {
	err = global.GVA_DB.Save(&fuelCustomerContract).Error
	return err
}

// GetFuelCustomerContract 根据id获取FuelCustomerContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContractService *FuelCustomerContractService) GetFuelCustomerContract(id uint) (err error, fuelCustomerContract autocode.FuelCustomerContract) {
	err = global.GVA_DB.Where("id = ?", id).First(&fuelCustomerContract).Error
	return
}

// GetFuelCustomerContractInfoList 分页获取FuelCustomerContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (fuelCustomerContractService *FuelCustomerContractService) GetFuelCustomerContractInfoList(info autoCodeReq.FuelCustomerContractSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.FuelCustomerContract{})
	var fuelCustomerContracts []autocode.FuelCustomerContract
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ContractCode != "" {
		db = db.Where("`contract_code` LIKE ?", "%"+info.ContractCode+"%")
	}
	if info.Customer_id != nil {
		db = db.Where("`customer_id` = ?", info.Customer_id)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Preload("FuelCustomer").Offset(offset).Find(&fuelCustomerContracts).Error
	return err, fuelCustomerContracts, total
}
