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

type AirportApi struct {
}

var airportService = service.ServiceGroupApp.AutoCodeServiceGroup.AirportService


// CreateAirport 创建Airport
// @Tags 航油系统
// @Summary 创建Airport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Airport true "创建Airport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /airport/createAirport [post]
func (airportApi *AirportApi) CreateAirport(c *gin.Context) {
	var airport autocode.Airport
	_ = c.ShouldBindJSON(&airport)
	if err := airportService.CreateAirport(airport); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAirport 删除Airport
// @Tags 航油系统
// @Summary 删除Airport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Airport true "删除Airport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /airport/deleteAirport [delete]
func (airportApi *AirportApi) DeleteAirport(c *gin.Context) {
	var airport autocode.Airport
	_ = c.ShouldBindJSON(&airport)
	if err := airportService.DeleteAirport(airport); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAirportByIds 批量删除Airport
// @Tags 航油系统
// @Summary 批量删除Airport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Airport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /airport/deleteAirportByIds [delete]
func (airportApi *AirportApi) DeleteAirportByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := airportService.DeleteAirportByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAirport 更新Airport
// @Tags 航油系统
// @Summary 更新Airport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Airport true "更新Airport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /airport/updateAirport [put]
func (airportApi *AirportApi) UpdateAirport(c *gin.Context) {
	var airport autocode.Airport
	_ = c.ShouldBindJSON(&airport)
	if err := airportService.UpdateAirport(airport); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAirport 用id查询Airport
// @Tags 航油系统
// @Summary 用id查询Airport
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Airport true "用id查询Airport"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /airport/findAirport [get]
func (airportApi *AirportApi) FindAirport(c *gin.Context) {
	var airport autocode.Airport
	_ = c.ShouldBindQuery(&airport)
	if err, reairport := airportService.GetAirport(airport.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reairport": reairport}, c)
	}
}

// GetAirportList 分页获取Airport列表
// @Tags 航油系统
// @Summary 分页获取Airport列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.AirportSearch true "分页获取Airport列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /airport/getAirportList [get]
func (airportApi *AirportApi) GetAirportList(c *gin.Context) {
	var pageInfo autocodeReq.AirportSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := airportService.GetAirportInfoList(pageInfo); err != nil {
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
