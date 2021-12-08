package service

import (
	"github.com/jefferygeng/yj/server/service/autocode"
	"github.com/jefferygeng/yj/server/service/example"
	"github.com/jefferygeng/yj/server/service/frontend"
	"github.com/jefferygeng/yj/server/service/system"
)

type ServiceGroup struct {
	ExampleServiceGroup  example.ServiceGroup
	SystemServiceGroup   system.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	FrontendServiceGroup frontend.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
