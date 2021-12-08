package request

import (
	"github.com/jefferygeng/yj/server/model/autocode"
	"github.com/jefferygeng/yj/server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}