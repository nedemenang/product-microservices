package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nedemenang/product-microservices/product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//Create the handlers
	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductionValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.Addroduct)
	postRouter.Use(ph.MiddlewareProductionValidation)
	
	s := &http.Server{
		Addr:              ":9090",
		Handler:           sm,
		TLSConfig:         nil,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)


	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	//Shutdown is a function that allows you to gracefully shut down a server in case it is in the process of
	//processing a request. The above context means that it will give a grace of 30 seconds before forcibly
	//Shutting down a server
	s.Shutdown(tc)
}
