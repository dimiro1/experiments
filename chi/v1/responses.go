package v1

import (
	"fmt"
	"net/http"

	"github.com/dimiro1/experiments/chi/entities"
)

type links struct {
	Self string `json:"self"`
}

// TodoResponse represents a response of Todo version 1
type TodoResponse struct {
	entities.Todo
	Version int   `json:"version"`
	Links   links `json:"_links"`
}

// TodosResponse render a collection of todos
type TodosResponse []*TodoResponse

// Render Good place to add headers into the response
func (t *TodoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	t.Links.Self = fmt.Sprintf("http://example.com/v1/todo/%d", t.ID)
	return nil
}

// Render render each single item and also populate the Link header
func (t TodosResponse) Render(w http.ResponseWriter, r *http.Request) error {
	for _, tr := range t {
		if err := tr.Render(w, r); err != nil {
			return err
		}
	}
	w.Header().Set("Link", `<http://example.com/v1/todos?p=2> rel="text"`)
	return nil
}

// NewTodoResponse returns a TodoResponse renderer
func NewTodoResponse(t entities.Todo) *TodoResponse {
	return &TodoResponse{
		Todo:    t,
		Version: 1,
	}
}

// NewTodosResponse returns a TodosResponse renderer
func NewTodosResponse(todos []entities.Todo) TodosResponse {
	renderers := TodosResponse{}

	for _, t := range todos {
		renderers = append(renderers, NewTodoResponse(t))
	}

	return renderers
}
