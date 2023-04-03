package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	commonController "github.com/nivb52/go-branches-learning/05-todo-CRUD-add-data-validations/finish/src/common/controller"
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

type TodoController struct {
	commonController.HttpRouter
}

func (t TodoController) getTodos() {
	jsonBytes, err := json.Marshal(todos)
	if err != nil {
		t.InternalServerError()
		return
	}
	t.Res.WriteHeader(http.StatusOK)
	t.Res.Write(jsonBytes)
}

func (t TodoController) getTodoById() {
	todoId := getIdFromUrl(t.Req.URL.Path)

	if todoId == "" {
		t.NotFound()
		return
	}

	todoPtr := getTodoFromDB(todoId)
	if todoPtr != nil {
		jsonBytes, err := json.Marshal(*todoPtr)
		t.MakeSuccessResponse(jsonBytes, err)
		return
	}
	// this should be impossible
	t.NotFound()
	return
}

func (t TodoController) createTodo() {
	defer t.Req.Body.Close()
	body, err := io.ReadAll(t.Req.Body)
	if err != nil {
		// connection error
		t.InternalServerError()
		return
	}

	var newData Todo
	err = json.Unmarshal(body, &newData)
	if err != nil {
		t.InternalServerError()
		return
	}
	fmt.Println("- newData ", newData)

	createdTodo, validationError := createTodoInDB(newData)
	if validationError != nil {
		t.BedRequestServerError(validationError)
		return
	}

	jsonBytes, err := json.Marshal(createdTodo)
	t.MakeSuccessResponse(jsonBytes, err)
	return
}

func (t TodoController) updateTodo() {
	todoId := getIdFromUrl(t.Req.URL.Path)
	if todoId == "" {
		t.NotFound()
		return
	}
	defer t.Req.Body.Close()
	body, err := io.ReadAll(t.Req.Body)
	if err != nil {
		// connection error
		// link: https://stackoverflow.com/questions/71338019/why-is-response-body-in-golang-is-a-readcloser
		t.InternalServerError()
		return
	}

	var newData Todo
	err = json.Unmarshal(body, &newData)
	if err != nil {
		t.InternalServerError()
		return
	}

	updatedTodo, validationError := updateTodoInDB(todoId, newData)
	if validationError != nil {
		t.BedRequestServerError(validationError)
		return
	} else if updatedTodo == nil {
		t.NotFound()
		return
	}

	jsonBytes, err := json.Marshal(updatedTodo)
	fmt.Println("- newData ", newData)
	t.MakeSuccessResponse(jsonBytes, err)
	return
}

func (t TodoController) deleteTodo() {
	todoId := getIdFromUrl(t.Req.URL.Path)
	if todoId == "" {
		t.NotFound()
		return
	}

	newTodos := deleteTodoInDB(todoId)
	if newTodos == nil {
		t.NotFound()
		return
	}

	jsonBytes, err := json.Marshal(newTodos)
	t.MakeSuccessResponse(jsonBytes, err)
	return
}
