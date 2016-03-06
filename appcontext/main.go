// The MIT License (MIT)

// Copyright (c) 2016 Claudemiro

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Please see this article
// https://elithrar.github.io/article/custom-handlers-avoiding-globals/

package main

import (
	"fmt"
	"net/http"
)

// The context could have database connection, redis, etc
// Everything that should be global, must be here.
type appContext struct {
	Name    string
	Version int
}

type appHandler struct {
	*appContext
	Handler func(*appContext, http.ResponseWriter, *http.Request)
}

// ServeHTTP could have error handling, authentication
func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler(h.appContext, w, r)
}

func main() {
	c := &appContext{Name: "Application Name", Version: 1}

	http.Handle("/", appHandler{appContext: c, Handler: indexHandler})
	http.ListenAndServe(":8080", nil)
}

func indexHandler(a *appContext, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IndexHandler: name is %s and version is %d", a.Name, a.Version)
}
