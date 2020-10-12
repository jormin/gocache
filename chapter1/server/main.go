package main

import (
	"github.com/jormin/go-cache/chapter1/server/cache"
	"github.com/jormin/go-cache/chapter1/server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
