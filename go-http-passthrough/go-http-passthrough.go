package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
)

// By default version flag is set to 1 (refers to HTTP/1.1)
var httpVersion = flag.Int("version", 1, "HTTP version")

var backendhost = flag.String("host", "localhost", "Server host")
var backendport = flag.String("port", "9191", "Server port")
var backendpath = flag.String("path", "/hello/sayHello", "Server Request path")

var certFile = flag.String("certfile", "", "SSL certificate file")
var keyFile = flag.String("keyfile", "", "SSL certificate key file")

// By default the number of maximum concurrent streams per connection is set as 1000
var maxConcurrentStreams = flag.Int("maxstream", 1000, "HTTP/2 max concurrent streams")

var client = &http.Client{}

func main() {
	flag.Parse()

	// Create a pool with the server certificate since it is not signed
	// by a known CA
	caCert, err := ioutil.ReadFile(*certFile)
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with the certificate of the server
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	// Use the proper transport in the client
	switch *httpVersion {
	case 1:
		client.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
		var httpServer = http.Server{
			Addr: ":9090",
		}
		http.HandleFunc("/passthrough", reverseProxy)
		log.Printf("Go Pssthrough: { HTTPVersion = 1 }; serving on https://localhost:9090/passthrough")
		log.Fatal(httpServer.ListenAndServeTLS(*certFile, *keyFile))
	case 2:
		client.Transport = &http2.Transport{
			TLSClientConfig: tlsConfig,
		}
		var httpServer = http.Server{
			Addr: ":9090",
		}
		var http2Server = http2.Server{
			MaxConcurrentStreams: uint32(*maxConcurrentStreams),
		}
		_ = http2.ConfigureServer(&httpServer, &http2Server)
		http.HandleFunc("/passthrough", reverseProxy)
		log.Printf("Go Pssthrough: { HTTPVersion = 2, MaxStreams = %v }; serving on https://localhost:9090/passthrough", *maxConcurrentStreams)
		log.Fatal(httpServer.ListenAndServeTLS(*certFile, *keyFile))
	}
}

func reverseProxy(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
	backendReq, _ := http.NewRequest(req.Method, "https://"+*backendhost+":"+*backendport+*backendpath, req.Body)
	resp, _ := client.Do(backendReq)
	body, _ := ioutil.ReadAll(resp.Body)
	_, _ = w.Write(body)
}
