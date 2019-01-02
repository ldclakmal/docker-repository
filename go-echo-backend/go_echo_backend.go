package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/sayHello", echoHttp)
	log.Printf("Serving on http://localhost:9191/hello/sayHello")
	log.Fatal(http.ListenAndServe(":9191", nil))
}

func echoHttp(writer http.ResponseWriter, req *http.Request) {
	req.Write(writer)
}
