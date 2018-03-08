package router

import (
	"net/http"
	"sync"

	"fmt"
	"log"
	"time"

	"github.com/gorilla/mux"
)

// SubRouter - sub router struct
type SubRouter struct {
	Name        string
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// ModuleRouter - 模块路由
type ModuleRouter struct {
	URLPrefix  string
	SubRouters []SubRouter
}

// SetURLPrefix - Set prefix
func (mr *ModuleRouter) SetURLPrefix(urlPrefix string) {
	mr.URLPrefix = urlPrefix
}

// SetSubRouters - 设置子路由
func (mr *ModuleRouter) SetSubRouters(srs []SubRouter) {
	mr.SubRouters = srs
}

// AddSubRouter - 添加子路由
func (mr *ModuleRouter) AddSubRouter(sr SubRouter) {
	mr.SubRouters = append(mr.SubRouters, sr)
}

// BaseRouter - 基础路由
type BaseRouter struct {
	R             *mux.Router
	ModuleRouters []ModuleRouter
	MainRouters   []SubRouter
	HandleNum     int
	BROnce        sync.Once
}

// Initialization - initial Router
func (br *BaseRouter) Initialization() {
	br.BROnce.Do(func() {
		br.RegisterMainRouter()
		br.RegisterSubRouter()
	})
	time.Sleep(50 * time.Millisecond)
	log.Printf("INFO(%d): Echo handle number.", br.HandleNum)
}

// AddModuleRouter - 添加模块路由
func (br *BaseRouter) AddModuleRouter(mr ModuleRouter) {
	br.ModuleRouters = append(br.ModuleRouters, mr)
}

// RegisterSubRouter - 注册子路由
func (br *BaseRouter) RegisterSubRouter() {
	for _, mr := range br.ModuleRouters {
		var sr = br.R.PathPrefix(mr.URLPrefix).Subrouter()
		for _, smr := range mr.SubRouters {
			sr.Methods(smr.Methods...).Path(smr.Pattern).Name(smr.Name).Handler(smr.HandlerFunc)
			// router counter
			br.HandleNum++
			// print router info
			fmt.Println(smr.Methods, mr.URLPrefix+smr.Pattern)
		}
	}
}

// RegisterMainRouter - register main router
func (br *BaseRouter) RegisterMainRouter() {
	for _, mr := range br.MainRouters {
		br.R.HandleFunc(mr.Pattern, mr.HandlerFunc).Methods(mr.Methods...).Name(mr.Name)
		// router counter
		br.HandleNum++
		// print router info
		fmt.Println(mr.Methods, mr.Pattern)
	}
}
