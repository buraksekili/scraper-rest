package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/buraksekili/scraper-rest/handlers"
	gohandlers "github.com/gorilla/handlers"
)

func main() {
	logger := log.New(os.Stdout, "scraper-api", log.LstdFlags)
	infoHandler := handlers.GetNewInfo(logger)

	serverRouter := mux.NewRouter()

	getRouter := serverRouter.Methods(http.MethodPost).Subrouter()
	getRouter.HandleFunc("/images", infoHandler.ParseImages)

	// CORS handler
	hCORS := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	server := http.Server{
		Addr:         ":5000",
		Handler:      hCORS(serverRouter),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Println("the server is running on port 5000")
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Error while listening: %s\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	log.Println("SIGNAL: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
