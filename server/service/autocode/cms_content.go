package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type CmsContentService struct {
}

// CreateCmsContent 创建CmsContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsContentService *CmsContentService) CreateCmsContent(cmsContent autocode.CmsContent) (err error) {
	err = global.GVA_DB.Create(&cmsContent).Error
	return err
}

// DeleteCmsContent 删除CmsContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsContentService *CmsContentService)DeleteCmsContent(cmsContent autocode.CmsContent) (err error) {
	err = global.GVA_DB.Delete(&cmsContent).Error
	return err
}

// DeleteCmsContentByIds 批量删除CmsContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsContentService *CmsContentService)DeleteCmsContentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.CmsContent{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCmsContent 更新CmsContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsContentService *CmsContentService)UpdateCmsContent(cmsContent autocode.CmsContent) (err error) {
	err = global.GVA_DB.Save(&cmsContent).Error
	return err
}

// GetCmsContent 根据id获取CmsContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsContentService *CmsContentService)GetCmsContent(id uint) (err error, cmsContent autocode.CmsContent) {
	err = global.GVA_DB.Where("id = ?", id).First(&cmsContent).Error
	return
}

// GetCmsContentInfoList 分页获取CmsContent记录
// Author [piexlmax](https://github.com/piexlmax)
func (cmsContentService *CmsContentService)GetCmsContentInfoList(info autoCodeReq.CmsContentSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.CmsContent{})
    var cmsContents []autocode.CmsContent
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Title != "" {
        db = db.Where("`title` LIKE ?","%"+ info.Title+"%")
    }
    if info.Aliasname != "" {
        db = db.Where("`aliasname` = ?",info.Aliasname)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&cmsContents).Error
	return err, cmsContents, total
}
