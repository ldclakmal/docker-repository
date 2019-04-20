package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
)

// By default version flag is set to 1 (refers to HTTP/1.1)
var httpVersion = flag.Int("version", 1, "HTTP version")

var certFile = flag.String("certfile", "", "SSL certificate file")
var keyFile = flag.String("keyfile", "", "SSL certificate key file")

// By default the number of maximum concurrent streams per connection is set as 1000
var maxConcurrentStreams = flag.Int("maxstream", 1000, "HTTP/2 max concurrent streams")

func main() {
	flag.Parse()

	switch *httpVersion {
	case 1:
		log.Printf("Go Backend: { HTTPVersion = 1 }; serving on https://localhost:%s%s", 9191, "/hello/sayHello")
		httpBackend()
	case 2:
		log.Printf("Go Backend: { HTTPVersion = 2, MaxStreams = %v }; serving on https://localhost:%s%s", *maxConcurrentStreams, 9191, "/hello/sayHello")
		http2Backend()
	}
}

func httpBackend() {
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Fatal(http.ListenAndServeTLS(":9191", *certFile, *keyFile, nil))
}

func http2Backend() {
	var httpServer = http.Server{
		Addr: ":9191",
	}
	var http2Server = http2.Server{
		MaxConcurrentStreams: uint32(*maxConcurrentStreams),
	}
	_ = http2.ConfigureServer(&httpServer, &http2Server)
	http.HandleFunc("/hello/sayHello", echoPayload)
	log.Fatal(httpServer.ListenAndServeTLS(*certFile, *keyFile))
}

func echoPayload(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	defer req.Body.Close()
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprintf(w, "%s\n", string(contents))
}
