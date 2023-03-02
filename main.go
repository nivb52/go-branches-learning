package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("nnot implemented"))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

// ////////////////////////////////
// reposetory.go TODOS REPO
type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone string `json:"isDone"`
}

var todos = []todo{{ID: "1", Title: "Learn GO", IsDone: "false"}, {ID: "2", Title: "Learn REACT", IsDone: "true"}}

// ////////////////////////////////
// services.go TODOS SERVICE
func getTodos(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func getTodoById(w http.ResponseWriter, q *http.Request) {
	//id := q.RequestURI()
	w.Write([]byte("<h1 style='color: red;'>getTodo By Id</h1><h2></h2>"))
}
func createTodo(w http.ResponseWriter, q *http.Request) {
	w.Write([]byte("<h1 style='color: red;'>createTodo</h1><h2>Data</h2><p></p>"))
}
func updateTodo(w http.ResponseWriter, q *http.Request) {
	w.Write([]byte("<h1 style='color: red;'>updateTodo By Id</h1><h2></h2>"))
}
func deleteTodo(w http.ResponseWriter, q *http.Request) {
	w.Write([]byte("<h1 style='color: red;'>deleteTodo By Id</h1><h2></h2>"))
}

// ////////////////////////////////
// controllers.go TODOS CONTROLLER
var (
	listTodosRegex        = regexp.MustCompile(`^\/todos[\/]*$`)
	todoRegex             = regexp.MustCompile(`^\/todos\/(\d+)$`)
	createTodoRegex       = regexp.MustCompile(`^\/todos[\/]*$`)
	createBatchTodosRegex = regexp.MustCompile(`^\/todos\/batch[\/]*$`)
)

func todosController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listTodosRegex.MatchString(r.URL.Path):
		getTodos(w, r)
		return
	case r.Method == http.MethodGet && todoRegex.MatchString(r.URL.Path):
		getTodoById(w, r)
		return
	case r.Method == http.MethodDelete && todoRegex.MatchString(r.URL.Path):
		deleteTodo(w, r)
		return
	case (r.Method == http.MethodPost || r.Method == http.MethodPatch) && todoRegex.MatchString(r.URL.Path):
		updateTodo(w, r)
		return
	case (r.Method == http.MethodPost) && (createTodoRegex.MatchString(r.URL.Path)):
		createTodo(w, r)
		return
	case (r.Method == http.MethodPost) && (createBatchTodosRegex.MatchString(r.URL.Path)):
		notImplemented(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

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
	mux.HandleFunc("/todos", todosController)
	log.Println(fmt.Sprintf("Starting Server on port %s", "5000"))
	log.Fatal(http.ListenAndServe(":5000", mux))
}
