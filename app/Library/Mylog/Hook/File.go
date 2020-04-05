package Hook

import (
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"echo-firefly-web/app/Library"
	"echo-firefly-web/app/Library/Yamls"
	"time"
)

func ConfigDayFilesHttpLogger() log.Hook {
	//LogPath := "/data/log/riskapi"
	LogPath := Yamls.GetConf().Other.SITE_LOG_DIR + "/echo-firefly-web/http"
	AccessLogPath := LogPath + "/access.log"
	ErrorLogPath := LogPath + "/error.log"
	logWriter, err := rotatelogs.New(
		AccessLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(AccessLogPath), // 生成软链，指向最新日志文件
		//rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config http access logger error. %+v", errors.WithStack(err))
	}
	logErrorWriter, err := rotatelogs.New(
		ErrorLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(ErrorLogPath), // 生成软链，指向最新日志文件
		//rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config http error logger error. %+v", errors.WithStack(err))
	}
	writeMap := lfshook.WriterMap{
		log.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
		log.InfoLevel:  logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logErrorWriter,
		log.FatalLevel: logErrorWriter,
		log.PanicLevel: logErrorWriter,
	}
	customFormatter := new(log.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	customFormatter.DisableTimestamp = false                // 禁止显示时间
	lfHook := lfshook.NewHook(writeMap, customFormatter)
	return lfHook
}

func ConfigFilesHttpLogger() log.Hook {
	//LogPath := "/data1/logs/app/nginx/cron"
	LogPath := Yamls.GetConf().Other.SITE_LOG_DIR + "/echo-firefly-web/http"
	AccessLogPath := LogPath + "/access.log"
	ErrorLogPath := LogPath + "/error.log"

	logWriter, err := os.OpenFile(AccessLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Errorf("config http access logger error. %+v", errors.WithStack(err))
	}
	logErrorWriter, err := os.OpenFile(ErrorLogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Errorf("config http access logger error. %+v", errors.WithStack(err))
	}

	writeMap := lfshook.WriterMap{
		log.DebugLevel: logWriter,
		log.InfoLevel:  logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logErrorWriter,
		log.FatalLevel: logErrorWriter,
		log.PanicLevel: logErrorWriter,
	}
	customFormatter := new(log.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	customFormatter.DisableTimestamp = false                // 禁止显示时间
	lfHook := lfshook.NewHook(writeMap, customFormatter)
	return lfHook
}

func ConfigFilesLogger(file string, path string) log.Hook {
	if (file == "") {
		file = "monitor"
	}
	if path == "" {
		path = "monitor"
	}
	logPath := Yamls.GetConf().Other.SITE_LOG_DIR + "/echo-firefly-web" + "/" + path
	if Library.Exist(logPath) != true {
		os.MkdirAll(logPath, os.ModePerm)
	}
	filePath := logPath + "/" + file + ".log"
	logWriter, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Errorf("config http access logger error. %+v", errors.WithStack(err))
	}
	writeMap := lfshook.WriterMap{
		log.DebugLevel: logWriter,
		log.InfoLevel:  logWriter,
		log.WarnLevel:  logWriter,
		log.ErrorLevel: logWriter,
		log.FatalLevel: logWriter,
		log.PanicLevel: logWriter,
	}
	customFormatter := new(log.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05" // 时间格式
	customFormatter.DisableTimestamp = false                // 禁止显示时间
	lfHook := lfshook.NewHook(writeMap, customFormatter)
	return lfHook
}
