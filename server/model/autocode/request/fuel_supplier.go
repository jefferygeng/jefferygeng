package request

import (
	"github.com/jefferygeng/yj/server/model/autocode"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type SupplierSearch struct {
	autocode.Supplier
	request.PageInfo
}
