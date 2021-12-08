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

type FuelCustomerContractApi struct {
}

var fuelCustomerContractService = service.ServiceGroupApp.AutoCodeServiceGroup.FuelCustomerContractService

// CreateFuelCustomerContract 创建FuelCustomerContract
// @Tags 航油系统
// @Summary 创建FuelCustomerContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerContract true "创建FuelCustomerContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomerContract/createFuelCustomerContract [post]
func (fuelCustomerContractApi *FuelCustomerContractApi) CreateFuelCustomerContract(c *gin.Context) {
	var fuelCustomerContract autocode.FuelCustomerContract
	_ = c.ShouldBindJSON(&fuelCustomerContract)
	if err := fuelCustomerContractService.CreateFuelCustomerContract(fuelCustomerContract); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFuelCustomerContract 删除FuelCustomerContract
// @Tags 航油系统
// @Summary 删除FuelCustomerContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerContract true "删除FuelCustomerContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fuelCustomerContract/deleteFuelCustomerContract [delete]
func (fuelCustomerContractApi *FuelCustomerContractApi) DeleteFuelCustomerContract(c *gin.Context) {
	var fuelCustomerContract autocode.FuelCustomerContract
	_ = c.ShouldBindJSON(&fuelCustomerContract)
	if err := fuelCustomerContractService.DeleteFuelCustomerContract(fuelCustomerContract); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFuelCustomerContractByIds 批量删除FuelCustomerContract
// @Tags 航油系统
// @Summary 批量删除FuelCustomerContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FuelCustomerContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fuelCustomerContract/deleteFuelCustomerContractByIds [delete]
func (fuelCustomerContractApi *FuelCustomerContractApi) DeleteFuelCustomerContractByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := fuelCustomerContractService.DeleteFuelCustomerContractByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFuelCustomerContract 更新FuelCustomerContract
// @Tags 航油系统
// @Summary 更新FuelCustomerContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerContract true "更新FuelCustomerContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fuelCustomerContract/updateFuelCustomerContract [put]
func (fuelCustomerContractApi *FuelCustomerContractApi) UpdateFuelCustomerContract(c *gin.Context) {
	var fuelCustomerContract autocode.FuelCustomerContract
	_ = c.ShouldBindJSON(&fuelCustomerContract)
	if err := fuelCustomerContractService.UpdateFuelCustomerContract(fuelCustomerContract); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFuelCustomerContract 用id查询FuelCustomerContract
// @Tags 航油系统
// @Summary 用id查询FuelCustomerContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.FuelCustomerContract true "用id查询FuelCustomerContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fuelCustomerContract/findFuelCustomerContract [get]
func (fuelCustomerContractApi *FuelCustomerContractApi) FindFuelCustomerContract(c *gin.Context) {
	var fuelCustomerContract autocode.FuelCustomerContract
	_ = c.ShouldBindQuery(&fuelCustomerContract)
	if err, refuelCustomerContract := fuelCustomerContractService.GetFuelCustomerContract(fuelCustomerContract.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refuelCustomerContract": refuelCustomerContract}, c)
	}
}

// GetFuelCustomerContractList 分页获取FuelCustomerContract列表
// @Tags 航油系统
// @Summary 分页获取FuelCustomerContract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.FuelCustomerContractSearch true "分页获取FuelCustomerContract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomerContract/getFuelCustomerContractList [get]
func (fuelCustomerContractApi *FuelCustomerContractApi) GetFuelCustomerContractList(c *gin.Context) {
	var pageInfo autocodeReq.FuelCustomerContractSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := fuelCustomerContractService.GetFuelCustomerContractInfoList(pageInfo); err != nil {
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
