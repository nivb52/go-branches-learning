package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nivb52/go-branches-learning/05-todo-CRUD-add-data-validations/starter/src/todo"
)

//////////////////////////////////
// main.go SERVER

func hello(rw http.ResponseWriter, rq *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	rw.Write([]byte("<h1 style='color: red;'>Hello World</h1>"))
}

func liveness(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("live"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", liveness)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/todos", todo.TodosRouter)
	mux.HandleFunc("/todos/", todo.TodosRouter)
	log.Println(fmt.Sprintf("Starting Server on port %s", "5000"))
	log.Fatal(http.ListenAndServe(":5000", mux))
}
