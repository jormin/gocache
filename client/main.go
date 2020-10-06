package main

import (
	"flag"
	"fmt"
	"go-cache/cache-benchmark/cacheClient"
)

func main() {
	s := flag.String("s", "127.0.0.1", "cache server address")
	c := flag.String("c", "get", "command，could be get|set|del")
	k := flag.String("k", "", "key")
	v := flag.String("v", "", "value")
	flag.Parse()
	client := cacheClient.New("tcp", *s)
	cmd := &cacheClient.Cmd{
		Name:  *c,
		Key:   *k,
		Value: *v,
		Error: nil,
	}
	client.Run(cmd)
	if cmd.Error != nil {
		fmt.Println("error:", cmd.Error)
	} else {
		fmt.Println(cmd.Value)
	}
}
