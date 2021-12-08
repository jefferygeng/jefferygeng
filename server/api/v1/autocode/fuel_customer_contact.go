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

type FuelCustomerContactApi struct {
}

var fuelCustomerContactService = service.ServiceGroupApp.AutoCodeServiceGroup.FuelCustomerContactService


// CreateFuelCustomerContact 创建FuelCustomerContact
// @Tags 航油系统
// @Summary 创建FuelCustomerContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerContact true "创建FuelCustomerContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomerContact/createFuelCustomerContact [post]
func (fuelCustomerContactApi *FuelCustomerContactApi) CreateFuelCustomerContact(c *gin.Context) {
	var fuelCustomerContact autocode.FuelCustomerContact
	_ = c.ShouldBindJSON(&fuelCustomerContact)
	if err := fuelCustomerContactService.CreateFuelCustomerContact(fuelCustomerContact); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFuelCustomerContact 删除FuelCustomerContact
// @Tags 航油系统
// @Summary 删除FuelCustomerContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerContact true "删除FuelCustomerContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fuelCustomerContact/deleteFuelCustomerContact [delete]
func (fuelCustomerContactApi *FuelCustomerContactApi) DeleteFuelCustomerContact(c *gin.Context) {
	var fuelCustomerContact autocode.FuelCustomerContact
	_ = c.ShouldBindJSON(&fuelCustomerContact)
	if err := fuelCustomerContactService.DeleteFuelCustomerContact(fuelCustomerContact); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFuelCustomerContactByIds 批量删除FuelCustomerContact
// @Tags 航油系统
// @Summary 批量删除FuelCustomerContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FuelCustomerContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fuelCustomerContact/deleteFuelCustomerContactByIds [delete]
func (fuelCustomerContactApi *FuelCustomerContactApi) DeleteFuelCustomerContactByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := fuelCustomerContactService.DeleteFuelCustomerContactByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFuelCustomerContact 更新FuelCustomerContact
// @Tags 航油系统
// @Summary 更新FuelCustomerContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.FuelCustomerContact true "更新FuelCustomerContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fuelCustomerContact/updateFuelCustomerContact [put]
func (fuelCustomerContactApi *FuelCustomerContactApi) UpdateFuelCustomerContact(c *gin.Context) {
	var fuelCustomerContact autocode.FuelCustomerContact
	_ = c.ShouldBindJSON(&fuelCustomerContact)
	if err := fuelCustomerContactService.UpdateFuelCustomerContact(fuelCustomerContact); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFuelCustomerContact 用id查询FuelCustomerContact
// @Tags 航油系统
// @Summary 用id查询FuelCustomerContact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.FuelCustomerContact true "用id查询FuelCustomerContact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fuelCustomerContact/findFuelCustomerContact [get]
func (fuelCustomerContactApi *FuelCustomerContactApi) FindFuelCustomerContact(c *gin.Context) {
	var fuelCustomerContact autocode.FuelCustomerContact
	_ = c.ShouldBindQuery(&fuelCustomerContact)
	if err, refuelCustomerContact := fuelCustomerContactService.GetFuelCustomerContact(fuelCustomerContact.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refuelCustomerContact": refuelCustomerContact}, c)
	}
}

// GetFuelCustomerContactList 分页获取FuelCustomerContact列表
// @Tags 航油系统
// @Summary 分页获取FuelCustomerContact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.FuelCustomerContactSearch true "分页获取FuelCustomerContact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fuelCustomerContact/getFuelCustomerContactList [get]
func (fuelCustomerContactApi *FuelCustomerContactApi) GetFuelCustomerContactList(c *gin.Context) {
	var pageInfo autocodeReq.FuelCustomerContactSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := fuelCustomerContactService.GetFuelCustomerContactInfoList(pageInfo); err != nil {
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
