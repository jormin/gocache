# About

基于 Tcp 的内存缓存服务

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

3. 性能测试

    测试机器需要安装 [Redis](https://redis.io)，从测试结果看：
    
    - 写：Redis 性能大约在自建的 Http 缓存服务的 7-8 倍左右。
    - 读：Redis 性能大约在自建的 Http 缓存服务的 11-12 倍左右。

    ```shell
    # 编译
    ➜  go-cache git:(master) cd cache-benchmark
    ➜  server git:(master) ✗ go build
    # 运行测试脚本
    ➜  go-cache git:(master) cd chapter1/test
    ➜  test git:(master) ✗ ./benchmark.sh
    [cp cache-benchmark to /usr/bin/]
    [cache-set] [cache-benchmark -type http -n 10000 -r 10000 -t set]
    type is http
    server is 127.0.0.1
    total 10000 requests
    data size is 1000
    we have 1 connections
    operation is set
    keyspacelen is 10000
    pipeline length is 1
    0 records get
    0 records miss
    10000 records set
    2.099862 seconds total
    99% requests < 1 ms
    99% requests < 2 ms
    99% requests < 3 ms
    99% requests < 4 ms
    100% requests < 5 ms
    207 usec average for each request
    throughput is 4.762218 MB/s
    rps is 4762.218454
    
    [cache-get] [cache-benchmark -type http -n 10000 -r 10000 -t get]
    type is http
    server is 127.0.0.1
    total 10000 requests
    data size is 1000
    we have 1 connections
    operation is get
    keyspacelen is 10000
    pipeline length is 1
    6349 records get
    3651 records miss
    0 records set
    3.572612 seconds total
    99% requests < 1 ms
    99% requests < 2 ms
    99% requests < 3 ms
    99% requests < 5 ms
    99% requests < 6 ms
    99% requests < 7 ms
    99% requests < 730 ms
    100% requests < 1167 ms
    346 usec average for each request
    throughput is 1.777131 MB/s
    rps is 2799.072486
    
    [redis-set,get] [redis-benchmark -c 1 -n 10000 -d 1000 -t set,get -r 10000]
    ====== SET ======
      10000 requests completed in 0.30 seconds
      1 parallel clients
      1000 bytes payload
      keep alive: 1
      host configuration "save": 3600 1 300 100 60 10000
      host configuration "appendonly": no
      multi-thread: no
    
    99.94% <= 0.1 milliseconds
    100.00% <= 1.5 milliseconds
    33444.82 requests per second
    
    ====== GET ======
      10000 requests completed in 0.27 seconds
      1 parallel clients
      1000 bytes payload
      keep alive: 1
      host configuration "save": 3600 1 300 100 60 10000
      host configuration "appendonly": no
      multi-thread: no
    
    99.96% <= 0.1 milliseconds
    100.00% <= 0.1 milliseconds
    36496.35 requests per second
    
    
    
    [remove cache-benchmark from /usr/bin/]
    
    ```