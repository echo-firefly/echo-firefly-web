package Hook

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	_ "github.com/zbindenren/logrus_mail"
	"echo-firefly-web/app/Library/Mylog/dao/logrusMail"
)

func ConfigMailLogger() log.Hook {
	hook, err := logrusMail.NewMailHook("RiskApi")
	if err != nil {
		log.Errorf("config mail logger error. %+v", errors.WithStack(err))
	}
	return hook
}
