package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

var logger = initLogger()

func initLogger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	logger.SetReportCaller(true)
	return logger
}

func Info(args ...interface{}) {
	logger.Log(logrus.InfoLevel, args...)
}

func Infoln(args ...interface{}) {
	logger.Logln(logrus.InfoLevel, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Logf(logrus.InfoLevel, format, args...)
}
