package v1

import (
	"server/api/v1/autocode"
	"server/api/v1/example"
	"server/api/v1/system"
)

type ApiGroup struct {
	ExampleApiGroup  example.ApiGroup
	SystemApiGroup   system.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
