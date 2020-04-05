package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"echo-firefly-web/app/Controllers"
	"echo-firefly-web/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper) {

	//定义全局控制器变量
	var C Controllers.Base

	b.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "欢迎使用echo-firefly-web框架")
	})
	g := b.Group("/test")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "这是测试主页面")
	})
	//加载到controller
	g.GET("/test", C.Test.GetData).Name = "test"
	g.GET("/api", C.Test.TestApi)
	g.GET("/redis", C.Test.Redis)


}
