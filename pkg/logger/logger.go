package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)

	// Optional: use JSON format
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Optional: set log level (DebugLevel, InfoLevel, WarnLevel, ErrorLevel)
	Log.SetLevel(logrus.InfoLevel)
}
