package autocode

import (
	"github.com/gin-gonic/gin"
	"github.com/jefferygeng/yj/server/global"
	"github.com/jefferygeng/yj/server/model/autocode"
	autocodeReq "github.com/jefferygeng/yj/server/model/autocode/request"
	"github.com/jefferygeng/yj/server/model/common/request"
	"github.com/jefferygeng/yj/server/model/common/response"
	"github.com/jefferygeng/yj/server/service"
	"go.uber.org/zap"
)

type SupplierApi struct {
}

var supplierService = service.ServiceGroupApp.AutoCodeServiceGroup.SupplierService


// CreateSupplier 创建Supplier
// @Tags 航油系统
// @Summary 创建Supplier
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Supplier true "创建Supplier"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /supplier/createSupplier [post]
func (supplierApi *SupplierApi) CreateSupplier(c *gin.Context) {
	var supplier autocode.Supplier
	_ = c.ShouldBindJSON(&supplier)
	if err := supplierService.CreateSupplier(supplier); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSupplier 删除Supplier
// @Tags 航油系统
// @Summary 删除Supplier
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Supplier true "删除Supplier"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /supplier/deleteSupplier [delete]
func (supplierApi *SupplierApi) DeleteSupplier(c *gin.Context) {
	var supplier autocode.Supplier
	_ = c.ShouldBindJSON(&supplier)
	if err := supplierService.DeleteSupplier(supplier); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSupplierByIds 批量删除Supplier
// @Tags 航油系统
// @Summary 批量删除Supplier
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Supplier"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /supplier/deleteSupplierByIds [delete]
func (supplierApi *SupplierApi) DeleteSupplierByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := supplierService.DeleteSupplierByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSupplier 更新Supplier
// @Tags 航油系统
// @Summary 更新Supplier
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Supplier true "更新Supplier"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /supplier/updateSupplier [put]
func (supplierApi *SupplierApi) UpdateSupplier(c *gin.Context) {
	var supplier autocode.Supplier
	_ = c.ShouldBindJSON(&supplier)
	if err := supplierService.UpdateSupplier(supplier); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSupplier 用id查询Supplier
// @Tags 航油系统
// @Summary 用id查询Supplier
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Supplier true "用id查询Supplier"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /supplier/findSupplier [get]
func (supplierApi *SupplierApi) FindSupplier(c *gin.Context) {
	var supplier autocode.Supplier
	_ = c.ShouldBindQuery(&supplier)
	if err, resupplier := supplierService.GetSupplier(supplier.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resupplier": resupplier}, c)
	}
}

// GetSupplierList 分页获取Supplier列表
// @Tags 航油系统
// @Summary 分页获取Supplier列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.SupplierSearch true "分页获取Supplier列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /supplier/getSupplierList [get]
func (supplierApi *SupplierApi) GetSupplierList(c *gin.Context) {
	var pageInfo autocodeReq.SupplierSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := supplierService.GetSupplierInfoList(pageInfo); err != nil {
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
