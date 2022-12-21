package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

func handleWannaBeTested(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("test me!\n")

	writer.WriteHeader(200)
	_, err := writer.Write([]byte("test me completely"))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var (
		port int
	)

	flag.IntVar(&port, "port", 3333, "enter port")
	flag.Parse()

	router := chi.NewRouter()
	router.Get("/test", handleWannaBeTested)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("server started at port - %v\n", port)

	<-done
	log.Print("server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		//extra handling here
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	log.Print("server exited properly")
}
