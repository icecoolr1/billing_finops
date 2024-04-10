package utils

import "github.com/sirupsen/logrus"

var Log = logrus.New()

func init() {
	Log.SetFormatter(&logrus.TextFormatter{})
	Log.SetLevel(logrus.InfoLevel)
	Log.SetReportCaller(true)
	Log.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:03:04",
	}
}
