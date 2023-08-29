package main

import (
	"log"
	"net"
	"net/http"

	"github.com/Jenny4831/hello/greeting" // import greeting the path is like module/package name
)

func main() {
	http.HandleFunc("/greeting", greeting.Greeting)

	log.Println("Starting server....")

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(listener, nil)
}
