package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/akhilesharora/flightpath/internal/api/handler"
)

func main() {
	svc := handler.New()

	// Register exit function to gracefully stop http server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case <-c:
			log.Println("stopping server")
			if err := svc.Close(); err != nil {
				log.Fatal(err)
			}
			log.Println("done")
		}
	}()

	url := net.JoinHostPort("0.0.0.0", "8080")

	log.Printf("starting server: %s\n", url)
	if err := svc.ListenAndServe(url); err != nil {
		log.Println("server stopped")
		os.Exit(1)
	}
}
