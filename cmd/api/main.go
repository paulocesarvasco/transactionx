package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"transactionx/internal/database"
	"transactionx/internal/handler"
	"transactionx/internal/router"
	"transactionx/internal/service"
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

	db := database.NewSQLiteClient()
	service := service.NewService(db)
	r := router.InstanceRoutes(handler.NewHandler(service))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
