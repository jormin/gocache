package main

import (
	"go-cache/chapter1/server/cache"
	"go-cache/chapter2/server/tcp"
)

func main() {
	c := cache.New("inmemory")
	tcp.New(c).Listen()
}
