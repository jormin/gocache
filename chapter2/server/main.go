package main

import (
	"github.com/jormin/go-cache/chapter1/server/cache"
	"github.com/jormin/go-cache/chapter2/server/http"
	"github.com/jormin/go-cache/chapter2/server/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
