package main

import (
	"context"
	"fmt"
	"log"
	"microservice/handellers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "Api", log.LstdFlags)

	logger.Println("Hellow")
	sm := http.NewServeMux()

	hellowHandeller := handellers.NewHellowHandller(logger)
	sm.Handle("/hellow", hellowHandeller)

	server := &http.Server{
		Addr:         ":7000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		server.ListenAndServe()

	}()
	sigCham := make(chan os.Signal)

	signal.Notify(sigCham, os.Interrupt)
	signal.Notify(sigCham, os.Kill)

	sig := <-sigCham

	logger.Println("Recieved The signal to shutdoen server ", sig)
	contextBackground, _ := context.WithTimeout(context.Background(), 60*time.Second)
	fmt.Println("Terminating Server")
	server.Shutdown(contextBackground)

}