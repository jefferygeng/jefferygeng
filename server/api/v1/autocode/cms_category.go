package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CmsCategoryApi struct {
}

var cmsCategoryService = service.ServiceGroupApp.AutoCodeServiceGroup.CmsCategoryService


// CreateCmsCategory 创建CmsCategory
// @Tags CmsCategory
// @Summary 创建CmsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsCategory true "创建CmsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCategory/createCmsCategory [post]
func (cmsCategoryApi *CmsCategoryApi) CreateCmsCategory(c *gin.Context) {
	var cmsCategory autocode.CmsCategory
	_ = c.ShouldBindJSON(&cmsCategory)
	if err := cmsCategoryService.CreateCmsCategory(cmsCategory); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsCategory 删除CmsCategory
// @Tags CmsCategory
// @Summary 删除CmsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsCategory true "删除CmsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsCategory/deleteCmsCategory [delete]
func (cmsCategoryApi *CmsCategoryApi) DeleteCmsCategory(c *gin.Context) {
	var cmsCategory autocode.CmsCategory
	_ = c.ShouldBindJSON(&cmsCategory)
	if err := cmsCategoryService.DeleteCmsCategory(cmsCategory); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsCategoryByIds 批量删除CmsCategory
// @Tags CmsCategory
// @Summary 批量删除CmsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsCategory/deleteCmsCategoryByIds [delete]
func (cmsCategoryApi *CmsCategoryApi) DeleteCmsCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cmsCategoryService.DeleteCmsCategoryByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsCategory 更新CmsCategory
// @Tags CmsCategory
// @Summary 更新CmsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsCategory true "更新CmsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsCategory/updateCmsCategory [put]
func (cmsCategoryApi *CmsCategoryApi) UpdateCmsCategory(c *gin.Context) {
	var cmsCategory autocode.CmsCategory
	_ = c.ShouldBindJSON(&cmsCategory)
	if err := cmsCategoryService.UpdateCmsCategory(cmsCategory); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsCategory 用id查询CmsCategory
// @Tags CmsCategory
// @Summary 用id查询CmsCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CmsCategory true "用id查询CmsCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsCategory/findCmsCategory [get]
func (cmsCategoryApi *CmsCategoryApi) FindCmsCategory(c *gin.Context) {
	var cmsCategory autocode.CmsCategory
	_ = c.ShouldBindQuery(&cmsCategory)
	if err, recmsCategory := cmsCategoryService.GetCmsCategory(cmsCategory.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsCategory": recmsCategory}, c)
	}
}

// GetCmsCategoryList 分页获取CmsCategory列表
// @Tags CmsCategory
// @Summary 分页获取CmsCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CmsCategorySearch true "分页获取CmsCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsCategory/getCmsCategoryList [get]
func (cmsCategoryApi *CmsCategoryApi) GetCmsCategoryList(c *gin.Context) {
	var pageInfo autocodeReq.CmsCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cmsCategoryService.GetCmsCategoryInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
