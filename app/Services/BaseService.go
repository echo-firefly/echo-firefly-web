package Services

import (
	"echo-firefly-web/app/Models"
)

//services层 主要是对业务层的一些逻辑封装
type Base struct {
	Test TestService
}
//关联model层
type TestService struct {
	Models Models.Base
}
