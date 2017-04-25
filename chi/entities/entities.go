package entities

// Todo represents a single Todo Item
type Todo struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
