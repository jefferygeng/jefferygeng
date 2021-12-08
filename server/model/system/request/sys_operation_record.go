package request

import (
	"github.com/jefferygeng/yj/server/model/common/request"
	"github.com/jefferygeng/yj/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
