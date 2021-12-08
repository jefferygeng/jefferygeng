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

type FuelSupplierContactApi struct {
}

var fuelSupplierContactService = service.ServiceGroupApp.AutoCodeServiceGroup.FuelSupplierContactService


// CreateFuelSupplierContact 创建FuelSupplierContact
// @Tags 航油系统
// @Summary 创建FuelSupplierContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelSupplierContact true "创建FuelSupplierContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelSupplierContact/createFuelSupplierContact [post]
func (fuelSupplierContactApi *FuelSupplierContactApi) CreateFuelSupplierContact(c *gin.Context) {
	var fuelSupplierContact autocode.FuelSupplierContact
	_ = c.ShouldBindJSON(&fuelSupplierContact)
	if err := fuelSupplierContactService.CreateFuelSupplierContact(fuelSupplierContact); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFuelSupplierContact 删除FuelSupplierContact
// @Tags 航油系统
// @Summary 删除FuelSupplierContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelSupplierContact true "删除FuelSupplierContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fuelSupplierContact/deleteFuelSupplierContact [delete]
func (fuelSupplierContactApi *FuelSupplierContactApi) DeleteFuelSupplierContact(c *gin.Context) {
	var fuelSupplierContact autocode.FuelSupplierContact
	_ = c.ShouldBindJSON(&fuelSupplierContact)
	if err := fuelSupplierContactService.DeleteFuelSupplierContact(fuelSupplierContact); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFuelSupplierContactByIds 批量删除FuelSupplierContact
// @Tags 航油系统
// @Summary 批量删除FuelSupplierContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FuelSupplierContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fuelSupplierContact/deleteFuelSupplierContactByIds [delete]
func (fuelSupplierContactApi *FuelSupplierContactApi) DeleteFuelSupplierContactByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := fuelSupplierContactService.DeleteFuelSupplierContactByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFuelSupplierContact 更新FuelSupplierContact
// @Tags 航油系统
// @Summary 更新FuelSupplierContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelSupplierContact true "更新FuelSupplierContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fuelSupplierContact/updateFuelSupplierContact [put]
func (fuelSupplierContactApi *FuelSupplierContactApi) UpdateFuelSupplierContact(c *gin.Context) {
	var fuelSupplierContact autocode.FuelSupplierContact
	_ = c.ShouldBindJSON(&fuelSupplierContact)
	if err := fuelSupplierContactService.UpdateFuelSupplierContact(fuelSupplierContact); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFuelSupplierContact 用id查询FuelSupplierContact
// @Tags 航油系统
// @Summary 用id查询FuelSupplierContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.FuelSupplierContact true "用id查询FuelSupplierContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fuelSupplierContact/findFuelSupplierContact [get]
func (fuelSupplierContactApi *FuelSupplierContactApi) FindFuelSupplierContact(c *gin.Context) {
	var fuelSupplierContact autocode.FuelSupplierContact
	_ = c.ShouldBindQuery(&fuelSupplierContact)
	if err, refuelSupplierContact := fuelSupplierContactService.GetFuelSupplierContact(fuelSupplierContact.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refuelSupplierContact": refuelSupplierContact}, c)
	}
}

// GetFuelSupplierContactList 分页获取FuelSupplierContact列表
// @Tags 航油系统
// @Summary 分页获取FuelSupplierContact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.FuelSupplierContactSearch true "分页获取FuelSupplierContact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelSupplierContact/getFuelSupplierContactList [get]
func (fuelSupplierContactApi *FuelSupplierContactApi) GetFuelSupplierContactList(c *gin.Context) {
	var pageInfo autocodeReq.FuelSupplierContactSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := fuelSupplierContactService.GetFuelSupplierContactInfoList(pageInfo); err != nil {
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
