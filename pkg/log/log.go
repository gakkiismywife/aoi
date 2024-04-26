package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func InitLog() {
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)
	Logger.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2016-01-02 15:04:05"})
}
