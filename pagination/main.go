package main

import (
	"encoding/json"
	"fmt"

	"github.com/dimiro1/experiments/pagination/paging"
)

// Post is our pageable object
type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// PostsPage represents a single page, it contains Page metadata and the Post results
type PostsPage struct {
	paging.Page
	Results []Post `json:"results"`
}

// AllParameters parameters for the All function
type AllParameters struct {
	paging.Params
}

// PostsRepository interface that defines the PostsRepositoey
type PostsRepository interface {
	All(AllParameters) PostsPage
}

// DummyPostsRepository implements the PostsRepository
type DummyPostsRepository struct{}

// All this is just a super basic example
// Here we can get posts in a database, some external api etc
func (DummyPostsRepository) All(params AllParameters) PostsPage {
	return PostsPage{
		Page: paging.Page{
			IsFirst:      true,
			IsLast:       false,
			Count:        10,
			Current:      1,
			PreviousLink: "", // We are ignoring empty
			NextLink:     "http://localhost:8000/posts?page=2",
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
	r := DummyPostsRepository{}
	page := r.All(AllParameters{paging.Params{
		Offset: 0,
		Limit:  10,
	}})
	data, _ := json.MarshalIndent(page, "", "\t")
	fmt.Println(string(data))
}
