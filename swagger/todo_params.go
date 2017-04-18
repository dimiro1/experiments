package main

// swagger:parameters getTodos
type getTodosInput struct {
	// Page number for pagination
	//
	// in:query
	// minimum:0
	Page uint64 `json:"page"`

	// Filter todos
	//
	// in:query
	Done bool `json:"done"`
}

// swagger:parameters createTodo
type createTodoInput struct {
	// in:body
	Todo todo `json:"todo"`
}
