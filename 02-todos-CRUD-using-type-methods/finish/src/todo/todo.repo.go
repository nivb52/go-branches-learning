package todo

import (
	"fmt"
	"strconv"

	utils "github.com/nivb52/go-branches-learning/02-todos-CRUD-using-type-methods/finish/src/utils"
)

// ////////////////////////////////
// reposetory.go TODOS REPO
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone string `json:"isDone"`
}

var lastTodoId = 16

func generateTodoId() string {
	lastTodoId = lastTodoId + 1
	lastTodoIdAsStr := strconv.Itoa(lastTodoId)
	return lastTodoIdAsStr
}

var todos = []Todo{{ID: "1", Title: "Learn GO", IsDone: "false"}, {ID: "2", Title: "Learn REACT", IsDone: "true"}, {ID: "15", Title: "Learn Advanced Go", IsDone: "false"}}

// ////////////////////////////////
// services.go TODOS SERVICE
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
			todo.IsDone = newData.IsDone
			todo.Title = newData.Title
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
