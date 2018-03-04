// paging package contains common pagination structures and functions
package pagination

// Metadata is a struct that represents a Metadata in both database and HTTP APIs
type Metadata struct {
	IsFirst     bool   `json:"is_first"`
	IsLast      bool   `json:"is_last"`
	HasNext     bool   `json:"has_next"`
	HasPrevious bool   `json:"has_previous"`
	Page        uint64 `json:"page"`
	Total       uint64 `json:"total"`
}

// ByOffset options to paginate by offset
type ByOffset struct {
	Offset uint64 `form:"offset"`
	Limit  uint64 `form:"-"`
}

// ByPageNum options to paginate by page num
type ByPageNum struct {
	Page    uint64 `form:"page"`
	PerPage uint64 `form:"per_page"`
}
