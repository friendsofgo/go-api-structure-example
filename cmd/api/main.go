package main

import (
	"os"

	"flag"

	"net/http"

	"fmt"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/ubeep/go-api-structure-example/pkg/server"
)

func main() {

	var (
		port    = os.Getenv("UBEEP_API_PORT")
		httpAdr = ":" + port
		//inmemory = flag.Bool("inmem", false, "use in-memory repositories")
	)

	flag.Parse()
	logger := logger()

	srv := server.New(logger)
	errs := make(chan error, 2)
	go func() {
		logger.WithFields(log.Fields{"transport": "http", "address": httpAdr}).Info("listening")
		errs <- http.ListenAndServe(httpAdr, srv)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Fatal("terminated ", <-errs)
}

func logger() *log.Logger {

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
