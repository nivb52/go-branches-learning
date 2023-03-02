package main

import (
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

//////////////////////////////////
// services.go TODOS SERVICE 
func getTodos (rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("<h1 style='color: red;'>getTodos</h1>"))
}
func getTodoById (rw http.ResponseWriter, rq *http.Request) {
	//id := rq.RequestURI()
	rw.Write([]byte("<h1 style='color: red;'>getTodo By Id</h1><h2></h2>"))
}
func createTodo (rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("<h1 style='color: red;'>createTodo</h1><h2>Data</h2><p></p>"))
}
func updateTodo (rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("<h1 style='color: red;'>updateTodo By Id</h1><h2></h2>"))
}
func deleteTodo (rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("<h1 style='color: red;'>deleteTodo By Id</h1><h2></h2>"))
}


//////////////////////////////////
// controllers.go TODOS CONTROLLER
var (
    listTodosRegex   = regexp.MustCompile(`^\/todos[\/]*$`)
    todoRegex    = regexp.MustCompile(`^\/todos\/(\d+)$`)
	createTodoRegex  = regexp.MustCompile(`^\/todos[\/]*$`)
    createBatchTodosRegex = regexp.MustCompile(`^\/todos\/batch[\/]*$`)
)

func todosController(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "application/json")
	switch {
		case r.Method == http.MethodGet  && listTodosRegex.MatchString(r.URL.Path):
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
	mux.HandleFunc("/todos/", todosController)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/", liveness)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
