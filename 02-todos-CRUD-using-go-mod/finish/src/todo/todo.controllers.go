package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	commonController "github.com/nivb52/go-branches-learning/02-todos-CRUD-using-type-methods/finish/src/common"
)

// ////////////////////////////////
// controllers.go TODOS CONTROLLER
var (
	listTodosRegex        = regexp.MustCompile(`^\/todos([\/]{0,1})$`)
	createTodoRegex       = regexp.MustCompile(`^\/todos([\/]{0,1})$`)
	todoWithIdRegex       = regexp.MustCompile(`^\/todos\/(\d+)$`)
	todoNamedFieldIdRegex = regexp.MustCompile(`^\/todos\/(?P<ID>([a-z,1-9,A-Z]{1,22}))\/?$`)
	createBatchTodosRegex = regexp.MustCompile(`^\/todos\/batch[\/]*$`)
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		commonController.InternalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
	todoId := getIdFromUrl(r.URL.Path)
	if todoId == "" {
		commonController.NotFound(w, r)
		return
	}

	successResponseFunc := makeSuccessResponse(w, r)
	todo := getTodoFromDB(todoId)
	if todo != nil {
		successResponseFunc(todo)
		return
	}

	// this is impossible
	commonController.NotFound(w, r)
	return
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// connection error
		commonController.InternalServerError(w, r)
		return
	}

	var newData Todo
	err = json.Unmarshal(body, &newData)
	if err != nil {
		commonController.InternalServerError(w, r)
		return
	}
	fmt.Println("- newData ", newData)

	createdTodo := createTodoInDB(&newData)

	successResponseFunc := makeSuccessResponse(w, r)
	successResponseFunc(createdTodo)
	return
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := getIdFromUrl(r.URL.Path)
	if todoId == "" {
		commonController.NotFound(w, r)
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// connection error
		// link: https://stackoverflow.com/questions/71338019/why-is-response-body-in-golang-is-a-readcloser
		commonController.InternalServerError(w, r)
		return
	}
	// x := string(body)
	// fmt.Println("- body ", x)
	// TODO: verify fields

	var newData Todo
	err = json.Unmarshal(body, &newData)
	if err != nil {
		commonController.InternalServerError(w, r)
		return
	}
	fmt.Println("- newData ", newData)

	updatedTodo := updateTodoInDB(todoId, newData)
	if updatedTodo == nil {
		commonController.NotFound(w, r)
		return
	}

	successResponseFunc := makeSuccessResponse(w, r)
	successResponseFunc(updatedTodo)
	return
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := getIdFromUrl(r.URL.Path)
	if todoId == "" {
		commonController.NotFound(w, r)
		return
	}

	newTodos := deleteTodoInDB(todoId)
	if newTodos == nil {
		commonController.NotFound(w, r)
		return
	}

	successResponseFunc := makeBatchSuccessResponse(w, r)
	successResponseFunc(newTodos)
	return
}
