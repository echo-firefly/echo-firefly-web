package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"project/echo-firefly-web/app/Controllers"
	"project/echo-firefly-web/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper) {

	b.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "欢迎使用echo-firefly-web框架")
	})
	g := b.Group("/test")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "这是测试主页面")
	})
	//加载到controller
	g.GET("/test", Controllers.TestGetData).Name = "test"
	g.GET("/redis", Controllers.TestRedis)


}
