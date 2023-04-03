package todo

import (
	"fmt"
	"reflect"
	"strconv"

	utils "github.com/nivb52/go-branches-learning/05-todo-CRUD-add-data-validations/finish/src/utils"
)

// ////////////////////////////////
// model.go TODOS MODEL

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func NewTodo(title string, isDone bool) *Todo {
	return &Todo{
		ID:     generateTodoId(),
		Title:  title,
		IsDone: isDone,
	}
}

func (t *Todo) GetIsDone() bool {
	return t.IsDone
}

func (t *Todo) SetIsDone(value bool) {
	t.IsDone = value
}

func (t *Todo) GetTitle() string {
	return t.Title
}

func (t *Todo) SetTitle(value string) {

	t.Title = value
}

var lastTodoId = 15

func generateTodoId() string {
	lastTodoId = lastTodoId + 1
	lastTodoIdAsStr := strconv.Itoa(lastTodoId)
	return lastTodoIdAsStr
}

var todos = []Todo{{ID: "1", Title: "Learn GO", IsDone: false}, {ID: "2", Title: "Learn REACT", IsDone: true}, {ID: "15", Title: "Learn Advanced Go", IsDone: false}}

// ////////////////////////////////
// services.go TODOS SERVICE // or // reposetory.go TODOS REPO
func getTodoFromDB(todoId string) *Todo {
	for _, todo := range todos {
		if todoId == todo.ID {
			return &todo
		}
	}
	return nil
}

func deleteTodoInDB(todoId string) *[]Todo {
	for index, todo := range todos {
		if todoId == todo.ID {
			todos = utils.SlicesDeleteFast(todos, index)
			return &todos

		}
	}
	return nil
}

func updateTodoInDB(todoId string, newData Todo) *Todo {
	for index, todo := range todos {
		if todoId == todo.ID {
			if reflect.TypeOf(todo.IsDone).Kind() == reflect.Bool {
				todo.IsDone = newData.IsDone
			}
			if reflect.TypeOf(todo.Title).Kind() == reflect.String {
				todo.Title = newData.Title
			}

			fmt.Printf("update in db %v", todo)
			todos[index] = todo
			return &todo
		}
	}
	return nil
}

func createTodoInDB(newTodo *Todo) *Todo {
	todoId := generateTodoId()
	newTodo.ID = todoId
	// newTodo := todo{ Title: title, IsActive: isActive, ID: todoId}
	todos = append(todos, *newTodo)
	return newTodo
}
