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

type FuelCustomerEmailApi struct {
}

var fuelCustomerEmailService = service.ServiceGroupApp.AutoCodeServiceGroup.FuelCustomerEmailService


// CreateFuelCustomerEmail 创建FuelCustomerEmail
// @Tags 航油系统
// @Summary 创建FuelCustomerEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerEmail true "创建FuelCustomerEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomerEmail/createFuelCustomerEmail [post]
func (fuelCustomerEmailApi *FuelCustomerEmailApi) CreateFuelCustomerEmail(c *gin.Context) {
	var fuelCustomerEmail autocode.FuelCustomerEmail
	_ = c.ShouldBindJSON(&fuelCustomerEmail)
	if err := fuelCustomerEmailService.CreateFuelCustomerEmail(fuelCustomerEmail); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFuelCustomerEmail 删除FuelCustomerEmail
// @Tags 航油系统
// @Summary 删除FuelCustomerEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerEmail true "删除FuelCustomerEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fuelCustomerEmail/deleteFuelCustomerEmail [delete]
func (fuelCustomerEmailApi *FuelCustomerEmailApi) DeleteFuelCustomerEmail(c *gin.Context) {
	var fuelCustomerEmail autocode.FuelCustomerEmail
	_ = c.ShouldBindJSON(&fuelCustomerEmail)
	if err := fuelCustomerEmailService.DeleteFuelCustomerEmail(fuelCustomerEmail); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFuelCustomerEmailByIds 批量删除FuelCustomerEmail
// @Tags 航油系统
// @Summary 批量删除FuelCustomerEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FuelCustomerEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fuelCustomerEmail/deleteFuelCustomerEmailByIds [delete]
func (fuelCustomerEmailApi *FuelCustomerEmailApi) DeleteFuelCustomerEmailByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := fuelCustomerEmailService.DeleteFuelCustomerEmailByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFuelCustomerEmail 更新FuelCustomerEmail
// @Tags 航油系统
// @Summary 更新FuelCustomerEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerEmail true "更新FuelCustomerEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fuelCustomerEmail/updateFuelCustomerEmail [put]
func (fuelCustomerEmailApi *FuelCustomerEmailApi) UpdateFuelCustomerEmail(c *gin.Context) {
	var fuelCustomerEmail autocode.FuelCustomerEmail
	_ = c.ShouldBindJSON(&fuelCustomerEmail)
	if err := fuelCustomerEmailService.UpdateFuelCustomerEmail(fuelCustomerEmail); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFuelCustomerEmail 用id查询FuelCustomerEmail
// @Tags 航油系统
// @Summary 用id查询FuelCustomerEmail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.FuelCustomerEmail true "用id查询FuelCustomerEmail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fuelCustomerEmail/findFuelCustomerEmail [get]
func (fuelCustomerEmailApi *FuelCustomerEmailApi) FindFuelCustomerEmail(c *gin.Context) {
	var fuelCustomerEmail autocode.FuelCustomerEmail
	_ = c.ShouldBindQuery(&fuelCustomerEmail)
	if err, refuelCustomerEmail := fuelCustomerEmailService.GetFuelCustomerEmail(fuelCustomerEmail.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refuelCustomerEmail": refuelCustomerEmail}, c)
	}
}

// GetFuelCustomerEmailList 分页获取FuelCustomerEmail列表
// @Tags 航油系统
// @Summary 分页获取FuelCustomerEmail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.FuelCustomerEmailSearch true "分页获取FuelCustomerEmail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomerEmail/getFuelCustomerEmailList [get]
func (fuelCustomerEmailApi *FuelCustomerEmailApi) GetFuelCustomerEmailList(c *gin.Context) {
	var pageInfo autocodeReq.FuelCustomerEmailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := fuelCustomerEmailService.GetFuelCustomerEmailInfoList(pageInfo); err != nil {
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
