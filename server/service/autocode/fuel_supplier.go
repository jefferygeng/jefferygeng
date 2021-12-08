package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type SupplierService struct {
}

// CreateSupplier 创建Supplier记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService) CreateSupplier(supplier autocode.Supplier) (err error) {
	err = global.GVA_DB.Create(&supplier).Error
	return err
}

// DeleteSupplier 删除Supplier记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)DeleteSupplier(supplier autocode.Supplier) (err error) {
	err = global.GVA_DB.Delete(&supplier).Error
	return err
}

// DeleteSupplierByIds 批量删除Supplier记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)DeleteSupplierByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Supplier{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSupplier 更新Supplier记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)UpdateSupplier(supplier autocode.Supplier) (err error) {
	err = global.GVA_DB.Save(&supplier).Error
	return err
}

// GetSupplier 根据id获取Supplier记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)GetSupplier(id uint) (err error, supplier autocode.Supplier) {
	err = global.GVA_DB.Where("id = ?", id).First(&supplier).Error
	return
}

// GetSupplierInfoList 分页获取Supplier记录
// Author [piexlmax](https://github.com/piexlmax)
func (supplierService *SupplierService)GetSupplierInfoList(info autoCodeReq.SupplierSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Supplier{})
    var suppliers []autocode.Supplier
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Supplier_code != "" {
        db = db.Where("`supplier_code` LIKE ?","%"+ info.Supplier_code+"%")
    }
    if info.Supplier_cn != "" {
        db = db.Where("`supplier_cn` LIKE ?","%"+ info.Supplier_cn+"%")
    }
    if info.Supplier_en != "" {
        db = db.Where("`supplier_en` LIKE ?","%"+ info.Supplier_en+"%")
    }
    if info.Supplier_small != "" {
        db = db.Where("`supplier_small` LIKE ?","%"+ info.Supplier_small+"%")
    }
    if info.Source != nil {
        db = db.Where("`source` = ?",info.Source)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&suppliers).Error
	return err, suppliers, total
}
