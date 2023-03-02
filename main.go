package main

import (
	"log"
	"net/http"
)

func hello(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("<h1 style='color: red;'>Hello World</h1>"))
}

func liveness(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("live"))
}

func main() {

	http.HandleFunc("/", liveness)
	http.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
