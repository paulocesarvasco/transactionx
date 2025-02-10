package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"transactionx/internal/constants"
	"transactionx/internal/database"
	"transactionx/internal/exchange"
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

	db := database.NewPostgresClient()
	exchangeService := exchange.NewService(&http.Client{}, constants.TREASURY_API_URL)
	service := service.NewService(db, exchangeService)
	r := router.InstanceRoutes(handler.NewHandler(service))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
