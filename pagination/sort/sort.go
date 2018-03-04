package sort

type Dir string

const (
	Asc  Dir = "asc"
	Desc     = "desc"
)

type Sortable struct {
	Sort  string `form:"sort"`
	Order Dir    `form:"order"`
}
