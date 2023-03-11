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
	message := "not found"
	// w.Write([]byte(message))
	jsonBytes, err := json.Marshal("{message:" + message + "}")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.Write(jsonBytes)
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	message := "not implemented"
	// w.Write([]byte(message))
	jsonBytes, err := json.Marshal("{message:" + message + "}")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.Write(jsonBytes)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.WriteHeader(http.StatusNotImplemented)
	message := "internal server error"
	// w.Write([]byte(message))
	jsonBytes, err := json.Marshal("{message:" + message + "}")
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.Write(jsonBytes)
}

// ////////////////////////////////
// utils.go TODO UTILS
var (
	listTodosRegex        = regexp.MustCompile(`^\/todos[\/]*$`)
	todoRegex             = regexp.MustCompile(`^\/todos\/(\d+)$`)
	createTodoRegex       = regexp.MustCompile(`^\/todos[\/]*$`)
	createBatchTodosRegex = regexp.MustCompile(`^\/todos\/batch[\/]*$`)
	todoIdRegex           = regexp.MustCompile(`^\/todos\/(?P<ID>([a-z,1-9,A-Z]{1,22}))\/?$`)
)

func makeSuccessResponse(w http.ResponseWriter, r *http.Request) func(*todo) {
	return func(data *todo) {
		jsonBytes, err := json.Marshal(&data)
		if err != nil {
			internalServerError(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	}
}

func makeBatchSuccessResponse(w http.ResponseWriter, r *http.Request) func(*[]todo) {
	return func(data *[]todo) {
		jsonBytes, err := json.Marshal(&data)
		if err != nil {
			internalServerError(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
		return
	}
}

// ////////////////////////////////
// reposetory.go TODOS REPO
type todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone string `json:"isDone"`
}

var todos = []todo{{ID: "1", Title: "Learn GO", IsDone: "false"}, {ID: "2", Title: "Learn REACT", IsDone: "true"}, {ID: "15", Title: "Learn Advanced Go", IsDone: "false"}}

// ////////////////////////////////
// services.go TODOS SERVICE
func getTodoFromDB(todoId string) *todo {
	for _, todo := range todos {
		if todoId == todo.ID {
			return &todo
		}
	}
	return nil
}

func deleteTodoFromDB(todoId string) *[]todo {
	for index, todo := range todos {
		if todoId == todo.ID {
			todos = slicesDeleteFast(todos, index)
			return &todos

		}
	}
	return nil
}

// ////////////////////////////////
// controllers.go TODOS CONTROLLER

func getTodos(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
	todoId := getIdFromUrl(r.URL.Path)
	if todoId == "" {
		notFound(w, r)
		return
	}

	successResponseFunc := makeSuccessResponse(w, r)
	todo := getTodoFromDB(todoId)
	if todo != nil {
		successResponseFunc(todo)
		return
	}

	// this is impossible
	notFound(w, r)
	return
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style='color: red;'>createTodo</h1><h2>Data</h2><p></p>"))
}
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style='color: red;'>updateTodo By Id</h1><h2></h2>"))
}
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := getIdFromUrl(r.URL.Path)
	if todoId == "" {
		notFound(w, r)
		return
	}

	newTodos := deleteTodoFromDB(todoId)
	successResponseFunc := makeBatchSuccessResponse(w, r)
	successResponseFunc(newTodos)
	return
}


// ////////////////////////////////
// routers.go TODOS CONTROLLER

func todosRouter(w http.ResponseWriter, r *http.Request) {
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
	mux.HandleFunc("/todos/", todosRouter)
	log.Println(fmt.Sprintf("Starting Server on port %s", "5000"))
	log.Fatal(http.ListenAndServe(":5000", mux))
}


//////////////////////////////////
// utils.go UTILS

// general utils
func slicesDeleteFast(s []todo, index int) []todo {
	if index >= len(s) || index < 0 {
		return s;
	} 
	if index == len(s)-1 {
    	return s[:len(s)-1]
	}
    s[index] = s[len(s)-1]
    return s[:len(s)-1]
}


//controller utils
func getIdFromUrl(url string) string {
	matches := todoIdRegex.FindStringSubmatch(url)
	var todoId string
	for i, name := range todoIdRegex.SubexpNames() {
		if name == "ID" {
			todoId = matches[i]
		}
	}
	fmt.Printf("todoId: %s\n", todoId)

	return todoId
}