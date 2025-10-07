package config

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogrus() {
	Log = logrus.New()
	Log.SetLevel(logrus.TraceLevel)
}
