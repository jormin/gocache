package main

import (
	"go-cache/chapter1/server/cache"
	"go-cache/chapter2/server/http"
	"go-cache/chapter2/server/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
