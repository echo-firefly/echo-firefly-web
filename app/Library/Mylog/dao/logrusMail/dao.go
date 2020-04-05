package logrusMail

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"echo-firefly-web/app/Library"
)

const (
	format = "20060102 15:04:05"
)

type Hook struct {
	AppName string
}

// NewMailHook creates a hook to be added to an instance of logger.
func NewMailHook(appname string) (*Hook, error) {
	return &Hook{AppName: appname}, nil
}

// Fire is called when a log event is fired.
func (hook *Hook) Fire(entry *logrus.Entry) error {
	err := createMessage(entry, hook.AppName)
	return err
}

// Levels returns the available logging levels.
func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

func createMessage(entry *logrus.Entry, appname string) error {
	subject := appname + " - " + entry.Level.String()
	body := entry.Time.Format(format) + " - " + entry.Message
	fields, _ := json.MarshalIndent(entry.Data, "", "\t")
	contents := fmt.Sprintf("%s\r\n\r\n%s\r\n%s", subject, body, string(fields))
	return Library.SendMail(subject, contents, []string{"localhost@wood.com"}, "develop")
}
