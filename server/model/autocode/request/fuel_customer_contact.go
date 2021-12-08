package request

import (
	"github.com/jefferygeng/yj/server/model/autocode"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type FuelCustomerContactSearch struct {
	autocode.FuelCustomerContact
	request.PageInfo
}
