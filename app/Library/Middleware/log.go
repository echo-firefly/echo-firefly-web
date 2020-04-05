package Middleware

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"echo-firefly-web/app/Library/Mylog"
)

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		Mylog.MakeHttpLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	var (
		code   = http.StatusInternalServerError
		msg    interface{}
		logMsg string
	)

	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		msg = he.Message
	} else {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		//msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}
	logMsg = fmt.Sprintf("code:%d,error:%v", he.Code, he.Message)
	Mylog.MakeHttpLogEntry(c).Error(logMsg)
	err = c.JSON(code, msg)
}
