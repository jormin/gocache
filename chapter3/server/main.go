package main

import (
	"flag"
	"github.com/jormin/go-cache/chapter1/server/cache"
	"github.com/jormin/go-cache/chapter3/server/http"
	"github.com/jormin/go-cache/chapter3/server/tcp"
	"log"
)

func main()  {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}