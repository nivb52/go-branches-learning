package todo

import (
	"net/http"

	commonController "github.com/nivb52/go-branches-learning/03-todos-CRUD-using-type-methods/finish/src/common"
)

// ////////////////////////////////
// routers.go TODOS CONTROLLER

func TodosRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listTodosRegex.MatchString(r.URL.Path):
		getTodos(w, r)
		return
	case r.Method == http.MethodGet && todoWithIdRegex.MatchString(r.URL.Path):
		getTodoById(w, r)
		return
	case r.Method == http.MethodDelete && todoWithIdRegex.MatchString(r.URL.Path):
		deleteTodo(w, r)
		return
	case (r.Method == http.MethodPost) && (createTodoRegex.MatchString(r.URL.Path)):
		createTodo(w, r)
		return
	case (r.Method == http.MethodPost || r.Method == http.MethodPut) && todoNamedFieldIdRegex.MatchString(r.URL.Path):
		updateTodo(w, r)
		return

	case (r.Method == http.MethodPost) && (createBatchTodosRegex.MatchString(r.URL.Path)):
		commonController.NotImplemented(w, r)
		return
	default:
		commonController.NotFound(w, r)
		return
	}
}
