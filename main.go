package main

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"BuffetSalesManage/BuffetSalesManage.git/src"
	"BuffetSalesManage/BuffetSalesManage.git/model/mongo"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	// 初始化mongo连接
	mongo.Connect()
	defer mongo.CloseSession()

	// 初始化Index
	mongo.InitializeIndex()

	// 初始化路由
	src.BuffetSalesRouter.Initialization()
	//log.Fatal(
	//	graceful.RunWithErr(
	//		router.FlagHostPort(13003), 120*time.Second, handlers.LoggingHandler(os.Stdout, src.KangarooRouter.R),
	//	),
	//)

	s := &http.Server{
		Addr:    ":13011",
		Handler: handlers.LoggingHandler(os.Stdout, src.BuffetSalesRouter.R),
	}

	go func() {
		log.Println("server BuffetSales start...")
		log.Fatal(s.ListenAndServe())
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Println(<-ch)

	if err := s.Shutdown(context.Background()); err != nil {
		log.Println("shutdown error:", err)
	}

}
