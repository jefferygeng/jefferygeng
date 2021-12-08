package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FuelCustomerContractSearch struct{
    autocode.FuelCustomerContract
    request.PageInfo
}