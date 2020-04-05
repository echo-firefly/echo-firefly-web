package Controllers

import "echo-firefly-web/app/Services"

//定义控制器层
type Base struct {
	Test TestController
}

//加载services
type TestController struct {
	Services Services.Base
}