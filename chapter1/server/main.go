package main

import (
	"go-cache/chapter1/server/cache"
	"go-cache/chapter1/server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
