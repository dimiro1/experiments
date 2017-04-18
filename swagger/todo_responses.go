package main

// swagger:model
type todo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// swagger:response getTodosResponse
type getTodosResponse struct {
	// in:body
	Todos []todo
}

// swagger:response unprocessableEntityResponse
type unprocessableEntityResponse struct {
	// in:body
	Payload struct {
		Status int `json:"status"`
		// Errors for each field
		Errors []map[string][]string `json:"errors"`
	}
}

// swagger:response emptyResponse
type emptyResponse struct{}

// swagger:response genericError
type errorResponse struct {
	// in:body
	Payload struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
}
