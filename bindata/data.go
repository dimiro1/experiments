// +build !embed

package main

import "net/http"

var Assets = http.Dir("public")
