package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CmsCategoryService struct {
}

// CreateCmsCategory 创建CmsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCategoryService *CmsCategoryService) CreateCmsCategory(cmsCategory autocode.CmsCategory) (err error) {
	err = global.GVA_DB.Create(&cmsCategory).Error
	return err
}

// DeleteCmsCategory 删除CmsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCategoryService *CmsCategoryService)DeleteCmsCategory(cmsCategory autocode.CmsCategory) (err error) {
	err = global.GVA_DB.Delete(&cmsCategory).Error
	return err
}

// DeleteCmsCategoryByIds 批量删除CmsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCategoryService *CmsCategoryService)DeleteCmsCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CmsCategory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsCategory 更新CmsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCategoryService *CmsCategoryService)UpdateCmsCategory(cmsCategory autocode.CmsCategory) (err error) {
	err = global.GVA_DB.Save(&cmsCategory).Error
	return err
}

// GetCmsCategory 根据id获取CmsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCategoryService *CmsCategoryService)GetCmsCategory(id uint) (err error, cmsCategory autocode.CmsCategory) {
	err = global.GVA_DB.Where("id = ?", id).First(&cmsCategory).Error
	return
}

// GetCmsCategoryInfoList 分页获取CmsCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsCategoryService *CmsCategoryService)GetCmsCategoryInfoList(info autoCodeReq.CmsCategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CmsCategory{})
    var cmsCategorys []autocode.CmsCategory
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Categoryname != "" {
        db = db.Where("`categoryname` = ?",info.Categoryname)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&cmsCategorys).Error
	return err, cmsCategorys, total
}
