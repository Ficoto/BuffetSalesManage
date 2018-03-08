package main

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"gitlab.xinghuolive.com/Backend-Go/orca/base"
	"gitlab.xinghuolive.com/Backend-Go/orca/model/mongo"
	"gitlab.xinghuolive.com/Backend-Go/orca/src"
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	// 初始化配置信息
	base.InitialConfigParam()

	// 初始化mongo连接
	mongo.Connect()
	defer mongo.CloseSession()

	//初始化Index
	mongo.InitializeIndex()

	// 初始化路由
	src.OrcaRouter.Initialization()
	//log.Fatal(
	//	graceful.RunWithErr(
	//		router.FlagHostPort(13003), 120*time.Second, handlers.LoggingHandler(os.Stdout, src.KangarooRouter.R),
	//	),
	//)

	s := &http.Server{
		Addr:    ":13011",
		Handler: handlers.LoggingHandler(os.Stdout, src.OrcaRouter.R),
	}

	go func() {
		log.Println("server orca start...")
		log.Fatal(s.ListenAndServe())
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Println(<-ch)

	if err := s.Shutdown(context.Background()); err != nil {
		log.Println("shutdown error:", err)
	}

}
