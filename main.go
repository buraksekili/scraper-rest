package main

import (
	"context"
	"github.com/buraksekili/scraper-rest/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "scraper-api", log.LstdFlags)
	infoHandler := handlers.GetNewInfo(logger)

	mux := http.NewServeMux()
	mux.Handle("/", infoHandler)

	server := http.Server{
		Addr: ":3000",
		Handler: mux,
		ErrorLog: logger,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		log.Println("the server is running on port 3000")
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
