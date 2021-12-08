package request

import (
	"github.com/jefferygeng/yj/server/model/common/request"
	"github.com/jefferygeng/yj/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
