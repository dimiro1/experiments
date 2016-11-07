// Example of best practices
// See: https://github.com/danielgtaylor/aglio
// See: https://github.com/apiaryio/dredd
// See: https://elithrar.github.io/article/http-handler-error-handling-revisited/
// See: https://text.sourcegraph.com/google-i-o-talk-building-sourcegraph-a-large-scale-code-search-cross-reference-engine-in-go-1f911b78a82e#.287nox7sm
// See: http://www.gorillatoolkit.org/pkg/
package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var defaultPerPage = 10

// Product represents a Product
type Product struct {
	ID   int64  `db:"id"   json:"id"`
	Name string `db:"name" json:"name"`
}

// SearchOptions is used to pass Search and Pagination parameters for respositories
type SearchOptions struct {
	PageOptions
	Q string `schema:"q"`
}

// DefaultSearchOptions returns a SearchOptions with its defaults
func DefaultSearchOptions() SearchOptions {
	return SearchOptions{
		PageOptions: DefaultPageOptions(),
	}
}

// PageOptions is used to pass Pagination parameters for respositories
type PageOptions struct {
	PerPage int `schema:"-"`
	Page    int `schema:"page"`
}

// Limit returns the PerPage number
func (p PageOptions) Limit() int { return p.PerPage }

// Offset returns the offset for pagination
func (p PageOptions) Offset() int { return (p.Page - 1) * p.PerPage }

// DefaultPageOptions returns a PageOptions with its defaults
func DefaultPageOptions() PageOptions {
	return PageOptions{
		PerPage: defaultPerPage,
	}
}

// InventoryRepository is an repository of inventory
type InventoryRepository interface {
	// All must return all from the database
	All(PageOptions) ([]Product, error)

	// Search, search products
	Search(SearchOptions) ([]Product, error)
}

// DatabaseInventoryRepository is a InventoryRepository baked by a SQL database
type DatabaseInventoryRepository struct {
	DB *sqlx.DB
}

// All must return all from the database using a SQL database
func (d *DatabaseInventoryRepository) All(options PageOptions) ([]Product, error) {
	products := []Product{}

	err := d.DB.Select(&products, "SELECT * FROM Products LIMIT ? OFFSET ?", options.Limit(), options.Offset())

	return products, err
}

// Search products using a SQL Database
func (d *DatabaseInventoryRepository) Search(options SearchOptions) ([]Product, error) {
	products := []Product{}

	err := d.DB.Select(&products, "SELECT * FROM Products WHERE name LIKE  '%' || ? || '%' LIMIT ? OFFSET ?",
		options.Q,
		options.Limit(), options.Offset(),
	)

	return products, err
}

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string { return se.Err.Error() }

// Status Returns our HTTP status code.
func (se StatusError) Status() int { return se.Code }

// ErrorHTTPHandler Special http handler that could returns errors when something goes wrong
type ErrorHTTPHandler func(*Env, http.ResponseWriter, *http.Request) error

// The Handler struct that wraps our handlers
type Handler struct {
	*Env
	H ErrorHTTPHandler
}

// Env application environment
type Env struct {
	inventory InventoryRepository
	decoder   *schema.Decoder
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)

	if err != nil {
		switch e := err.(type) {
		case Error:
			w.WriteHeader(e.Status())
			http.Error(w, e.Error(), e.Status())
		default:
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// ProductsHandler handle http requests on /
func ProductsHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	options := DefaultPageOptions()

	if err := env.decoder.Decode(&options, r.URL.Query()); err != nil {
		return StatusError{http.StatusBadRequest, err}
	}

	if options.Page <= 0 {
		options.Page = 1
	}

	products, err := env.inventory.All(options)

	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		return err
	}

	return nil
}

// SearchHandler handle http requests on /search
func SearchHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	options := DefaultSearchOptions()

	if err := env.decoder.Decode(&options, r.URL.Query()); err != nil {
		return StatusError{http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))}
	}

	if options.Page <= 0 {
		options.Page = 1
	}

	products, err := env.inventory.Search(options)

	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		return err
	}

	return nil
}

func main() {
	db := sqlx.MustOpen("sqlite3", ":memory:")

	db.MustExec(`CREATE TABLE Products (
                        id INTEGER PRIMARY KEY, 
                        name TEXT
                );
                INSERT INTO Products VALUES (1, 'TV');
                INSERT INTO Products VALUES (2, 'Microwave');`)

	inventory := &DatabaseInventoryRepository{db}

	env := &Env{
		inventory: inventory,
		decoder:   schema.NewDecoder(),
	}

	r := mux.NewRouter()

	r.Handle("/products", handlers.LoggingHandler(os.Stdout, Handler{env, ProductsHandler}))
	r.Handle("/products/search", handlers.LoggingHandler(os.Stdout, Handler{env, SearchHandler}))
	r.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))

	http.ListenAndServe(":8080", r)
}
