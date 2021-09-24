package router

import (
	"server/router/autocode"
	"server/router/example"
	"server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
