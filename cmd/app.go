package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Logger initialize logger configuration for the application
func Logger() *log.Logger {

	logger := log.New()
	logger.SetFormatter(loggerFmt())
	logger.SetOutput(os.Stdout)

	lvl, err := log.ParseLevel(os.Getenv("UBEEP_LOG_LEVEL"))
	if err != nil {
		logger.SetLevel(log.ErrorLevel)
		return logger
	}

	logger.SetLevel(lvl)

	return logger
}

func loggerFmt() log.Formatter {

	if os.Getenv("UBEEP_ENV") == "development" {
		formatter := new(log.TextFormatter)
		formatter.FullTimestamp = true
		return formatter
	} else {
		formatter := new(log.JSONFormatter)
		return formatter
	}

}
