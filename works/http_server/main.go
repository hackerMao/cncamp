package main

import (
	"fmt"
	"http_server/routers"
	"log"
	"net/http"
	"time"
)

func main() {
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	r := routers.InitRouter()
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    time.Second * 60,
		WriteTimeout:   time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server on port %d", port)
	}
}
