package request

import (
	"github.com/jefferygeng/yj/server/model/autocode"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type FuelSupplierContactSearch struct {
	autocode.FuelSupplierContact
	request.PageInfo
}
