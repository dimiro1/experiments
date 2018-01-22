// paging package contains common pagination structures and functions
package paging

// Page is a struct that represents a Page in both database and HTTP APIs
type Page struct {
	IsFirst      bool   `json:"is_first"`
	IsLast       bool   `json:"is_last"`
	NextLink     string `json:"next_link,omitempty"`
	PreviousLink string `json:"previous_link,omitempty"`
	Current      int64  `json:"current"`
	Count        int64  `json:"count"`
}

// Params groups common pagination parameters
type Params struct {
	Offset uint64
	Limit  uint64
}