package main

import (
	"os"
	"os/signal"

	"github.com/Sirupsen/logrus"
	"BuffetSalesManage/BuffetSalesManage.git/model/mongo"
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
