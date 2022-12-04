package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Printf("server: %s %v\n", r.Method, string(body))
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", 3333),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(10000 * time.Millisecond)
}
