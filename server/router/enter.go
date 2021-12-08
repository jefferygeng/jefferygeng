package router

import (
	"github.com/jefferygeng/yj/server/router/autocode"
	"github.com/jefferygeng/yj/server/router/example"
	"github.com/jefferygeng/yj/server/router/frontend"
	"github.com/jefferygeng/yj/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
	Frontend frontend.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
