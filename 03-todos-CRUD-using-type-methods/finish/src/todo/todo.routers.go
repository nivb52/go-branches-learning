package todo

import (
	"net/http"
)

// ////////////////////////////////
// routers.go TODOS CONTROLLER

func TodosRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var serv TodoController
	serv.Res = w
	serv.Req = r
	// serv := commonController.HttpContext{Writer: w, Req: *r}
	switch {
	case r.Method == http.MethodGet && listTodosRegex.MatchString(r.URL.Path):
		serv.getTodos()
		return
	case r.Method == http.MethodGet && todoWithIdRegex.MatchString(r.URL.Path):
		serv.getTodoById()
		return
	case r.Method == http.MethodDelete && todoWithIdRegex.MatchString(r.URL.Path):
		serv.deleteTodo()
		return
	case (r.Method == http.MethodPost) && (createTodoRegex.MatchString(r.URL.Path)):
		serv.createTodo()
		return
	case (r.Method == http.MethodPost || r.Method == http.MethodPut) && todoNamedFieldIdRegex.MatchString(r.URL.Path):
		serv.updateTodo()
		return

	case (r.Method == http.MethodPost) && (createBatchTodosRegex.MatchString(r.URL.Path)):
		serv.NotImplemented()
		return
	default:
		serv.NotFound()
		return
	}
}
