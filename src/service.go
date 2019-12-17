/**
 * 业务模块
 **/

package main

import (
	"common/git"
	"common/mylog"
	"config"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * 包装ServeHTTP并打印访问日志
 **/
func myServeHTTP(route http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() { AccessLog(r) }()

		r.Body = http.MaxBytesReader(w, r.Body, 8192)
		w.Header().Set("Content-Type", "application/json")

		route.ServeHTTP(w, r)
	})
}

/**
 *  Start Harbour Service
 **/
func StartWebService(cfg *config.Service) error {
	if cfg == nil {
		mylog.LogError("Config [Service] Error.")
		return errors.New("Config [Service] Error.")
	}

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(MyNotFoundHandler)

	/** 第一级目录 **/
	//v1r := r.PathPrefix("/Harbour").Subrouter()

	/** 业务操作 **/
	r.HandleFunc("/captain", captainHandler)
	r.HandleFunc("/captain:{op:.*}", captainOperateHandler)
	r.HandleFunc("/upload", uploadHandler)

	httpServer := http.Server{
		Handler:      myServeHTTP(r),
		Addr:         ":" + strconv.Itoa(cfg.ListenPort),
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		mylog.LogError("http.ListenAndServe: %v", err)
		return err
	}
	return nil
}

func Start(cfg *config.Service) error {

	go my_timer() // 启动定时器
	return StartWebService(cfg)
}

func my_timer() {
	//git.GitClone("https://github.com/yycaptain/adapter-go.git")
	ticker := time.NewTicker(time.Minute * 1) // 1分钟的ticker

	for range ticker.C {
		AutoCheck() // 自动检查
	}
}

func AccessLog(r *http.Request) {
	// TODO
	mylog.LogAccess("access log.")
}

func AutoCheck() {
	mylog.LogInfo("auto check log.")

}

func captainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			// TODO
		}
	case "POST":
		{

		}
	}
}

func captainOperateHandler(w http.ResponseWriter, r *http.Request) {
	op := mux.Vars(r)["op"]
	if strings.EqualFold(op, "cat") {
		// TODO
	} else if strings.EqualFold(op, "dog") {
		// TODO
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello."))
}
