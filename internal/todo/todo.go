package todo

import (
	"errors"
	"strings"
)

type Item struct {
	Task   string
	Status string
}

type Service struct {
	todos []Item
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

func (svc *Service) Add(todo string) error {
	// add validation since todo needs to be unique
	// t is accessing the item of the todos
	for _, t := range svc.todos {
		if t.Task == todo {
			return errors.New("todo is not unique")
		}
	}
	svc.todos = append(svc.todos, Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	})

	// if unique then we'll have no error
	return nil
}

func (svc *Service) GetAll() []Item {
	return svc.todos
}

func (svc *Service) Search(query string) []string {
	var results []string
	for _, todo := range svc.todos {
		if strings.Contains(strings.ToLower(todo.Task), strings.ToLower(query)) {
			results = append(results, todo.Task)
		}
	}
	return results
}
