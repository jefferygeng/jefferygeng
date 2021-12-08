package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FuelCustomerApi struct {
}

var fuelCustomerService = service.ServiceGroupApp.AutoCodeServiceGroup.FuelCustomerService


// CreateFuelCustomer 创建FuelCustomer
// @Tags 航油系统
// @Summary 创建FuelCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomer true "创建FuelCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomer/createFuelCustomer [post]
func (fuelCustomerApi *FuelCustomerApi) CreateFuelCustomer(c *gin.Context) {
	var fuelCustomer autocode.FuelCustomer
	_ = c.ShouldBindJSON(&fuelCustomer)
	if err := fuelCustomerService.CreateFuelCustomer(fuelCustomer); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFuelCustomer 删除FuelCustomer
// @Tags 航油系统
// @Summary 删除FuelCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomer true "删除FuelCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fuelCustomer/deleteFuelCustomer [delete]
func (fuelCustomerApi *FuelCustomerApi) DeleteFuelCustomer(c *gin.Context) {
	var fuelCustomer autocode.FuelCustomer
	_ = c.ShouldBindJSON(&fuelCustomer)
	if err := fuelCustomerService.DeleteFuelCustomer(fuelCustomer); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFuelCustomerByIds 批量删除FuelCustomer
// @Tags 航油系统
// @Summary 批量删除FuelCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FuelCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fuelCustomer/deleteFuelCustomerByIds [delete]
func (fuelCustomerApi *FuelCustomerApi) DeleteFuelCustomerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := fuelCustomerService.DeleteFuelCustomerByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFuelCustomer 更新FuelCustomer
// @Tags 航油系统
// @Summary 更新FuelCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomer true "更新FuelCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fuelCustomer/updateFuelCustomer [put]
func (fuelCustomerApi *FuelCustomerApi) UpdateFuelCustomer(c *gin.Context) {
	var fuelCustomer autocode.FuelCustomer
	_ = c.ShouldBindJSON(&fuelCustomer)
	if err := fuelCustomerService.UpdateFuelCustomer(fuelCustomer); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFuelCustomer 用id查询FuelCustomer
// @Tags 航油系统
// @Summary 用id查询FuelCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.FuelCustomer true "用id查询FuelCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fuelCustomer/findFuelCustomer [get]
func (fuelCustomerApi *FuelCustomerApi) FindFuelCustomer(c *gin.Context) {
	var fuelCustomer autocode.FuelCustomer
	_ = c.ShouldBindQuery(&fuelCustomer)
	if err, refuelCustomer := fuelCustomerService.GetFuelCustomer(fuelCustomer.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refuelCustomer": refuelCustomer}, c)
	}
}

// GetFuelCustomerList 分页获取FuelCustomer列表
// @Tags 航油系统
// @Summary 分页获取FuelCustomer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.FuelCustomerSearch true "分页获取FuelCustomer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomer/getFuelCustomerList [get]
func (fuelCustomerApi *FuelCustomerApi) GetFuelCustomerList(c *gin.Context) {
	var pageInfo autocodeReq.FuelCustomerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := fuelCustomerService.GetFuelCustomerInfoList(pageInfo); err != nil {
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
