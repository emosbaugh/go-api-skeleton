package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// initLogging sets the logging level and formatter being used.
func initLogging() {
	lvl, err := logrus.ParseLevel(viper.GetString("log-level"))
	if err != nil {
		logrus.
			WithField("err", err).
			WithField("log-level", viper.GetString("log-level")).
			Warning("invalid log level, default info")
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)

	var fmt logrus.Formatter
	switch viper.GetString("log-format") {
	case "text":
		fmt = new(logrus.TextFormatter)
	case "json":
		fmt = new(logrus.JSONFormatter)
	default:
		logrus.WithField("log-format", viper.GetString("log-format")).Warning("invalid log format, default text")
		fmt = new(logrus.TextFormatter)
	}
	logrus.SetFormatter(fmt)
}
