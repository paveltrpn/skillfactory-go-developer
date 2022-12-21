package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"skillfactory-go-developer/pkg/citiesdb"
)

var (
	reqPort int = 3333
)

func makeAndSendReq(data []byte, method string, port int, url string) error {
	reqUrl := fmt.Sprintf("http://localhost:%d%v", port, url)
	req, err := http.NewRequest(method, reqUrl, bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Charset":      {"utf-8"},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(res.Status)
	resBody, _ := io.ReadAll(res.Body)
	fmt.Println(string(resBody))

	return nil
}

func requestTest(port int) {
	err := makeAndSendReq([]byte{}, http.MethodGet, port, "/test")
	if err != nil {
		log.Fatalf("requestUpdateAge(): error sending request - %v\n", err)
	}
}

func requestAddCity(info citiesdb.CityInfo)
func main() {
	var (
		action string
	)

	flag.IntVar(&reqPort, "port", 3333, "enter port")
	flag.Parse()

	fmt.Printf("request sending to port - %v\n", reqPort)

	for {
		fmt.Printf("choose action - test\n")
		fmt.Scan(&action)

		switch action {
		case "test":
			fmt.Printf("test\n")

			requestTest(reqPort)

		default:
			fmt.Printf("unknown action - %v\n", action)
		}
	}
}
