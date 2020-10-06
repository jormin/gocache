# About

基于 Http 的内存缓存服务

1. 编译运行

```shell
➜  go-cache git:(master) cd chapter1/server 
➜  server git:(master) ✗ go build
➜  server git:(master) ✗ ./server 
2020/10/06 17:51:17 inmemory ready to serve
2020/10/06 17:51:36 http: superfluous response.WriteHeader call from go-cache/chapter1/server/http.(*cacheHandler).ServeHTTP (cache.go:67)
...
```

2. 请求测试

```shell
➜  go-cache git:(master) cd chapter1/test
 ➜  test git:(master) ✗ ./test.sh 
 [get stat] [curl 127.0.0.1:9090/status/]
 {"Count":0,"KeySize":0,"ValueSize":0}
 
 [set cache] [curl -v 127.0.0.1:9090/cache/testkey -XPUT -dtestvalue]
 *   Trying 127.0.0.1...
 * TCP_NODELAY set
 * Connected to 127.0.0.1 (127.0.0.1) port 9090 (#0)
 > PUT /cache/testkey HTTP/1.1
 > Host: 127.0.0.1:9090
 > User-Agent: curl/7.64.1
 > Accept: */*
 > Content-Length: 9
 > Content-Type: application/x-www-form-urlencoded
 > 
 * upload completely sent off: 9 out of 9 bytes
 < HTTP/1.1 200 OK
 < Date: Tue, 06 Oct 2020 10:13:08 GMT
 < Content-Length: 52
 < Content-Type: text/plain; charset=utf-8
 < 
 * Connection #0 to host 127.0.0.1 left intact
 {"code":0,"message":"put cache success","data":null}
 
 * Closing connection 0
 [get cache] [curl 127.0.0.1:9090/cache/testkey]
 {"code":0,"message":"get cache success","data":"testvalue"}
 
 [get stat] [curl 127.0.0.1:9090/status/]
 {"Count":1,"KeySize":7,"ValueSize":9}
 
 [delete cache] [curl 127.0.0.1:9090/cache/testkey -XDELETE]
 {"code":0,"message":"delete cache success","data":null}
 
 [get stat] [curl 127.0.0.1:9090/status/]
 {"Count":0,"KeySize":0,"ValueSize":0}
 
```