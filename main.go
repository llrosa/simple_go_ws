package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got %s request at %s\n", r.Method, r.RequestURI)
	fmt.Printf("headers: %v\n", r.Header)
	fmt.Printf("remoteAddr %s\nhost %s\n", r.RemoteAddr, r.Host)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := "<h1>Ha! You Got me</h1>"
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got %s request at %s\n", r.Method, r.RequestURI)
	fmt.Printf("headers: %v\n", r.Header)
	fmt.Printf("remoteAddr %s\nhost %s\n", r.RemoteAddr, r.Host)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	data := "Hello"
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func main() {
	port := 9090
	if p := os.Getenv("PORT"); p != "" {
		if i, err := strconv.Atoi(p); err == nil {
			port = i
		}
	}
	fmt.Printf("Server started at %d\n", port)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
