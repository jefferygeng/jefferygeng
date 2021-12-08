package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ZnFlightService struct {
}

// CreateZnFlight 创建ZnFlight记录
// Author [piexlmax](https://github.com/piexlmax)
func (znFlightService *ZnFlightService) CreateZnFlight(znFlight autocode.ZnFlight) (err error) {
	err = global.GVA_DB.Create(&znFlight).Error
	return err
}

// DeleteZnFlight 删除ZnFlight记录
// Author [piexlmax](https://github.com/piexlmax)
func (znFlightService *ZnFlightService)DeleteZnFlight(znFlight autocode.ZnFlight) (err error) {
	err = global.GVA_DB.Delete(&znFlight).Error
	return err
}

// DeleteZnFlightByIds 批量删除ZnFlight记录
// Author [piexlmax](https://github.com/piexlmax)
func (znFlightService *ZnFlightService)DeleteZnFlightByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ZnFlight{},"id in ?",ids.Ids).Error
	return err
}

// UpdateZnFlight 更新ZnFlight记录
// Author [piexlmax](https://github.com/piexlmax)
func (znFlightService *ZnFlightService)UpdateZnFlight(znFlight autocode.ZnFlight) (err error) {
	err = global.GVA_DB.Save(&znFlight).Error
	return err
}

// GetZnFlight 根据id获取ZnFlight记录
// Author [piexlmax](https://github.com/piexlmax)
func (znFlightService *ZnFlightService)GetZnFlight(id uint) (err error, znFlight autocode.ZnFlight) {
	err = global.GVA_DB.Where("id = ?", id).First(&znFlight).Error
	return
}

// GetZnFlightInfoList 分页获取ZnFlight记录
// Author [piexlmax](https://github.com/piexlmax)
func (znFlightService *ZnFlightService)GetZnFlightInfoList(info autoCodeReq.ZnFlightSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.ZnFlight{})
    var znFlights []autocode.ZnFlight
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&znFlights).Error
	return err, znFlights, total
}
