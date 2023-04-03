package todo

import (
	"errors"
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

func NewTodo(todo Todo) (Todo, error) {
	newTodo := Todo{
		ID: generateTodoId(),
	}
	err := newTodo.SetIsDone(todo.IsDone)
	if err != nil {
		return todo, err
	}
	err = newTodo.SetTitle(todo.Title)
	if err != nil {
		return todo, err
	}

	return newTodo, nil
}

func (t *Todo) GetIsDone() bool {
	return t.IsDone
}

func (t *Todo) SetIsDone(value bool) error {
	if reflect.TypeOf(value).Kind() != reflect.Bool {
		return errors.New("todo field isDone must be a boolean")
	}
	t.IsDone = value
	return nil
}

func (t *Todo) GetTitle() string {
	return t.Title
}

func (t *Todo) SetTitle(value string) error {
	if reflect.TypeOf(value).Kind() != reflect.String {
		return errors.New("todo field title must be a string")
	} else if len(value) < 3 {
		return errors.New("title is too short, please provide at least 3 chars")
	} else if len(value) > 35 {
		return errors.New("title is too long, please provide less than 35 chars")
	}
	t.Title = value
	return nil
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

func updateTodoInDB(todoId string, newData Todo) (*Todo, error) {
	for index, todo := range todos {
		if todoId == todo.ID {
			err := todo.SetIsDone(newData.IsDone)
			if err != nil {
				return &todo, err
			}
			err = todo.SetTitle(newData.Title)
			if err != nil {
				return &todo, err
			}
			fmt.Printf("success update in db %v", todo)
			todos[index] = todo
			return &todo, nil
		}
	}
	return nil, nil
}

func createTodoInDB(newData Todo) (*Todo, error) {
	newTodo, err := NewTodo(newData)

	if err != nil {
		return nil, err
	}
	todos = append(todos, newTodo)
	fmt.Printf("success created in db %v", newTodo)
	return &newTodo, nil
}
