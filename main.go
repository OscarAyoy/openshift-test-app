package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var hostName, _ = os.Hostname()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	log.Printf("Received request from %q\n", ip)
	fmt.Fprintf(w, "Hello from %s!", hostName)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
