package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"transactionx/internal/handler"
	"transactionx/internal/router"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func(c chan os.Signal) {
		<-c
		log.Println("Shutting down server")
		os.Exit(0)
	}(c)

	r := router.InstanceRoutes(handler.NewHandler())

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
