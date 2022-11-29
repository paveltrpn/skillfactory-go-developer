package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handler(signal os.Signal) {
	if signal == syscall.SIGINT {
		fmt.Printf("Got CTRL+C signal, closing!\n")
		os.Exit(0)
	}
}

func main() {
	sigchnl := make(chan os.Signal, 1)
	signal.Notify(sigchnl)

	go func() {
		for {
			s := <-sigchnl
			handler(s)
		}
	}()

	var nbr int

	for {
		fmt.Scan(&nbr)
		fmt.Printf("squared - %v\n", nbr*nbr)
	}
}
