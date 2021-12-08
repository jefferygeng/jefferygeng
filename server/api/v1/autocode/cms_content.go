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

type CmsContentApi struct {
}

var cmsContentService = service.ServiceGroupApp.AutoCodeServiceGroup.CmsContentService


// CreateCmsContent 创建CmsContent
// @Tags CmsContent
// @Summary 创建CmsContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsContent true "创建CmsContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsContent/createCmsContent [post]
func (cmsContentApi *CmsContentApi) CreateCmsContent(c *gin.Context) {
	var cmsContent autocode.CmsContent
	_ = c.ShouldBindJSON(&cmsContent)
	if err := cmsContentService.CreateCmsContent(cmsContent); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCmsContent 删除CmsContent
// @Tags CmsContent
// @Summary 删除CmsContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsContent true "删除CmsContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cmsContent/deleteCmsContent [delete]
func (cmsContentApi *CmsContentApi) DeleteCmsContent(c *gin.Context) {
	var cmsContent autocode.CmsContent
	_ = c.ShouldBindJSON(&cmsContent)
	if err := cmsContentService.DeleteCmsContent(cmsContent); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCmsContentByIds 批量删除CmsContent
// @Tags CmsContent
// @Summary 批量删除CmsContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CmsContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cmsContent/deleteCmsContentByIds [delete]
func (cmsContentApi *CmsContentApi) DeleteCmsContentByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := cmsContentService.DeleteCmsContentByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCmsContent 更新CmsContent
// @Tags CmsContent
// @Summary 更新CmsContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.CmsContent true "更新CmsContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cmsContent/updateCmsContent [put]
func (cmsContentApi *CmsContentApi) UpdateCmsContent(c *gin.Context) {
	var cmsContent autocode.CmsContent
	_ = c.ShouldBindJSON(&cmsContent)
	if err := cmsContentService.UpdateCmsContent(cmsContent); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCmsContent 用id查询CmsContent
// @Tags CmsContent
// @Summary 用id查询CmsContent
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.CmsContent true "用id查询CmsContent"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cmsContent/findCmsContent [get]
func (cmsContentApi *CmsContentApi) FindCmsContent(c *gin.Context) {
	var cmsContent autocode.CmsContent
	_ = c.ShouldBindQuery(&cmsContent)
	if err, recmsContent := cmsContentService.GetCmsContent(cmsContent.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recmsContent": recmsContent}, c)
	}
}

// GetCmsContentList 分页获取CmsContent列表
// @Tags CmsContent
// @Summary 分页获取CmsContent列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.CmsContentSearch true "分页获取CmsContent列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cmsContent/getCmsContentList [get]
func (cmsContentApi *CmsContentApi) GetCmsContentList(c *gin.Context) {
	var pageInfo autocodeReq.CmsContentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := cmsContentService.GetCmsContentInfoList(pageInfo); err != nil {
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
