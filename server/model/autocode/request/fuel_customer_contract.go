package request

import (
	"github.com/jefferygeng/yj/server/model/autocode"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type FuelCustomerContractSearch struct {
	autocode.FuelCustomerContract
	request.PageInfo
}
