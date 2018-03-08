package main

import (
	"os"
	"os/signal"

	"github.com/Sirupsen/logrus"
	"gitlab.xinghuolive.com/Backend-Go/orca/model/mongo"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)

	mongo.Connect()
	defer mongo.CloseSession()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
}
