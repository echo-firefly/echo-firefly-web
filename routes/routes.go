package routes

import (
	"github.com/labstack/echo"
	"net/http"
	"project/echo_web/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper) {

	b.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is a index")
	})
	g := b.Group("/test")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is a test")
	})


}
