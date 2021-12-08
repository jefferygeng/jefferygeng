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

type ZnFlightApi struct {
}

var znFlightService = service.ServiceGroupApp.AutoCodeServiceGroup.ZnFlightService


// CreateZnFlight 创建ZnFlight
// @Tags ZnFlight
// @Summary 创建ZnFlight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ZnFlight true "创建ZnFlight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /znFlight/createZnFlight [post]
func (znFlightApi *ZnFlightApi) CreateZnFlight(c *gin.Context) {
	var znFlight autocode.ZnFlight
	_ = c.ShouldBindJSON(&znFlight)
	if err := znFlightService.CreateZnFlight(znFlight); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteZnFlight 删除ZnFlight
// @Tags ZnFlight
// @Summary 删除ZnFlight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ZnFlight true "删除ZnFlight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /znFlight/deleteZnFlight [delete]
func (znFlightApi *ZnFlightApi) DeleteZnFlight(c *gin.Context) {
	var znFlight autocode.ZnFlight
	_ = c.ShouldBindJSON(&znFlight)
	if err := znFlightService.DeleteZnFlight(znFlight); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteZnFlightByIds 批量删除ZnFlight
// @Tags ZnFlight
// @Summary 批量删除ZnFlight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ZnFlight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /znFlight/deleteZnFlightByIds [delete]
func (znFlightApi *ZnFlightApi) DeleteZnFlightByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := znFlightService.DeleteZnFlightByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateZnFlight 更新ZnFlight
// @Tags ZnFlight
// @Summary 更新ZnFlight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ZnFlight true "更新ZnFlight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /znFlight/updateZnFlight [put]
func (znFlightApi *ZnFlightApi) UpdateZnFlight(c *gin.Context) {
	var znFlight autocode.ZnFlight
	_ = c.ShouldBindJSON(&znFlight)
	if err := znFlightService.UpdateZnFlight(znFlight); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindZnFlight 用id查询ZnFlight
// @Tags ZnFlight
// @Summary 用id查询ZnFlight
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ZnFlight true "用id查询ZnFlight"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /znFlight/findZnFlight [get]
func (znFlightApi *ZnFlightApi) FindZnFlight(c *gin.Context) {
	var znFlight autocode.ZnFlight
	_ = c.ShouldBindQuery(&znFlight)
	if err, reznFlight := znFlightService.GetZnFlight(znFlight.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reznFlight": reznFlight}, c)
	}
}

// GetZnFlightList 分页获取ZnFlight列表
// @Tags ZnFlight
// @Summary 分页获取ZnFlight列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ZnFlightSearch true "分页获取ZnFlight列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /znFlight/getZnFlightList [get]
func (znFlightApi *ZnFlightApi) GetZnFlightList(c *gin.Context) {
	var pageInfo autocodeReq.ZnFlightSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := znFlightService.GetZnFlightInfoList(pageInfo); err != nil {
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
