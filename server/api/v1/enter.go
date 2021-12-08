package v1

import (
	"github.com/jefferygeng/yj/server/api/v1/autocode"
	"github.com/jefferygeng/yj/server/api/v1/example"
	"github.com/jefferygeng/yj/server/api/v1/frontend"
	"github.com/jefferygeng/yj/server/api/v1/system"
)

type ApiGroup struct {
	ExampleApiGroup  example.ApiGroup
	SystemApiGroup   system.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
	FrontendApiGroup frontend.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
