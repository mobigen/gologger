package main

import (
	"fmt"

	formatter "github.com/mobigen/gologger"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Print("\n--- mobigen-formatter ---\n\n")
	printDemo(&formatter.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		ShowFields:      true,
	}, "mobigen-formatter")

	fmt.Print("\n--- default logrus formatter ---\n\n")
	printDemo(nil, "default logrus formatter")
}

func printDemo(f logrus.Formatter, title string) {
	l := logrus.New()

	l.SetLevel(logrus.DebugLevel)

	if f != nil {
		l.SetFormatter(f)
	}

	// enable/disable file/function name
	l.SetReportCaller(true)

	l.Infof("this is %v demo", title)

	lWebServer := l.WithField("component", "web-server")
	lWebServer.Info("starting...")

	lWebServerReq := lWebServer.WithFields(logrus.Fields{
		"req":   "GET /api/stats",
		"reqId": "#1",
	})

	lWebServerReq.Info("params: startYear=2048")
	lWebServerReq.Error("response: 400 Bad Request")

	lDbConnector := l.WithField("category", "db-connector")
	lDbConnector.Info("connecting to db on 10.10.10.13...")
	lDbConnector.Warn("connection took 10s")

	l.Info("demo end.")
}
