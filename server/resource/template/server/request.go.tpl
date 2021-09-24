package request

import (
	"server/model/autocode"
	"server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}