package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"skillfactory-go-developer/pkg/user"
)

var (
	reqUrl string = "http://localhost:3333/"
)

func main() {
	var (
		user user.User
		name string
		age  uint
	)

	for {
		fmt.Scanf("%v %v", &name, &age)

		user.Name = name
		user.Age = age

		req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader([]byte("hello!")))
		if err != nil {
			log.Fatalf("error making request - %v\n", err)
		}

		req.Header = http.Header{
			"Content-Type": {"application/json"},
			"Charset":      {"utf-8"},
		}

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			log.Fatalf("error sending request - %v\n", err)
		}
	}
}
