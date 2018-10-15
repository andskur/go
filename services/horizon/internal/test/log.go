package test

import (
	"github.com/andskur/go/support/log"
	"github.com/sirupsen/logrus"
)

var testLogger *log.Entry

func init() {
	testLogger = log.New()
	testLogger.Entry.Logger.Formatter.(*logrus.TextFormatter).DisableColors = true
	testLogger.Entry.Logger.Level = logrus.DebugLevel
}
