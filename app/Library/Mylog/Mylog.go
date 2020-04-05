package Mylog

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"echo-firefly-web/app/Library/Mylog/Hook"
	"time"
)

func MakeHttpLogEntry(c echo.Context) *log.Entry {
	var logObj = log.New()
	logObj.AddHook(Hook.ConfigFilesHttpLogger())
	logObj.AddHook(Hook.ConfigMailLogger())
	if c == nil {
		return logObj.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logObj.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func MakeLogEntry(c echo.Context, fields map[string]interface{}) *log.Entry {
	var logObj = log.New()
	logObj.AddHook(Hook.ConfigFilesHttpLogger())
	logObj.AddHook(Hook.ConfigMailLogger())
	if c == nil {
		fields["at"] = time.Now().Format("2006-01-02 15:04:05");
	} else {
		fields["at"] = time.Now().Format("2006-01-02 15:04:05");
		fields["method"] = c.Request().Method;
		fields["uri"] = c.Request().URL.String();
		fields["ip"] = c.Request().RemoteAddr;
	}
	return logObj.WithFields(fields)
}

func MakeLogger(file string, path string, c echo.Context) *log.Logger {
	logObj := log.New()
	logObj.AddHook(Hook.ConfigFilesLogger(file, path))
	logObj.AddHook(Hook.ConfigMailLogger())
	fields := log.Fields{}
	if c == nil {
		fields["at"] = time.Now().Format("2006-01-02 15:04:05");
	} else {
		fields["at"] = time.Now().Format("2006-01-02 15:04:05");
		fields["method"] = c.Request().Method;
		fields["uri"] = c.Request().URL.String();
		fields["ip"] = c.Request().RemoteAddr;
	}
	logObj.WithFields(fields)
	return logObj
}

func MakeLog() *log.Logger {
	return MakeLogger("", "", nil)
}
