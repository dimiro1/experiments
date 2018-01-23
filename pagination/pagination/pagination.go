// paging package contains common pagination structures and functions
package pagination

// Metadata is a struct that represents a Metadata in both database and HTTP APIs
type Metadata struct {
	IsFirst      bool   `json:"is_first"`
	IsLast       bool   `json:"is_last"`
	NextLink     string `json:"next_link,omitempty"`
	PreviousLink string `json:"previous_link,omitempty"`
	Current      int64  `json:"current"`
	Count        int64  `json:"count"`
}

// ByOffset options to paginate by offset
type ByOffset struct {
	Offset uint64 `schema:"offset"`
	Limit  uint64 `schema:"-"`
}

// ByPageNum options to paginate by page num
type ByPageNum struct {
	Page    uint64 `schema:"page"`
	PerPage uint64 `schema:"-"`
}
