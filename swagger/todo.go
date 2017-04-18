package main

import "net/http"

// swagger:route GET /todos todos getTodos
//
// List todos filtered by some parameters.
//
// 		Produces:
//			- application/json
//
//		Responses:
//			default: genericError
//			200: getTodosResponse
func getTodos(w http.ResponseWriter, r *http.Request) {}

// swagger:route POST /todos todos createTodo
//
// Create a new todo item.
//
//		Consumes:
//			- application/json
//
//		Produces:
//			- application/json
//
//		Responses:
//			default: genericError
//			422: unprocessableEntityResponse
//			201: emptyResponse
func createTodo(w http.ResponseWriter, r *http.Request) {}
