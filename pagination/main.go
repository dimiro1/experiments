package main

import (
	"net/http"

	"github.com/dimiro1/experiments/pagination/pagination"
	"github.com/dimiro1/experiments/pagination/sort"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Post is our pageable object
type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// PagedPosts represents a single page, it contains Metadata metadata and the Post results
type PagedPosts struct {
	pagination.Metadata
	Results []Post `json:"results"`
}

// FindAllOptions parameters for the FindAll function
type FindAllOptions struct {
	sort.Sortable
	pagination.ByPageNum
}

// PostsRepository interface that defines the repository with posts
type PostsRepository interface {
	FindAll(FindAllOptions) PagedPosts
}

// DummyPostsRepository implements the PostsRepository
type DummyPostsRepository struct{}

// FindAll this is just a super basic example
// Here we can get posts in a database, some external api etc
func (DummyPostsRepository) FindAll(options FindAllOptions) PagedPosts {
	return PagedPosts{
		Metadata: pagination.Metadata{
			IsFirst:     true,
			IsLast:      false,
			Total:       options.PerPage,
			Page:        options.Page,
			HasNext:     true,
			HasPrevious: false,
		},
		Results: []Post{
			{
				Title: "First Post",
				Body:  "Lorem ipsun dollor sit amet...",
			},
			{
				Title: "Second Post",
				Body:  "Lorem ipsun dollor sit amet...",
			},
		},
	}
}

func main() {
	var posts PostsRepository = DummyPostsRepository{}
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		var params FindAllOptions
		if err := ctx.BindQuery(&params); err != nil {
			ctx.Negotiate(http.StatusBadRequest, gin.Negotiate{Data: err, Offered: []string{binding.MIMEJSON, binding.MIMEXML}})
			return
		}

		page := posts.FindAll(params)

		ctx.Header("Link", `<http://localhost:8000?page=1>; rel="next"`)
		ctx.Negotiate(http.StatusOK, gin.Negotiate{Data: page, Offered: []string{binding.MIMEJSON, binding.MIMEXML}})
	})

	r.Run()
}
