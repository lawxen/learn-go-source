package main

import (
	"baby/models"
	"baby/routers"
	"net/http"
	"time"
)

func init() {
	models.Setup()
}

func main() {
	server := &http.Server{
		Addr:         ":8000",
		Handler:      routers.InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// 可以使用fvbock/endless替换http的ListenAndServe实现平滑重启
	server.ListenAndServe()
}
