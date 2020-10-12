# About

基于 Tcp 的内存缓存服务

1. 编译运行

    ```shell
    ➜  github.com/jormin/go-cache git:(master) cd chapter2/server 
    ➜  server git:(master) ✗ go build
    ➜  server git:(master) ✗ ./server 
    2020/10/07 00:04:25 inmemory ready to serve
    ...
    ```

2. 请求测试

    ```shell
    ➜  github.com/jormin/go-cache git:(master) cd chapter2/test
    ➜  test git:(master) ✗ ./test.sh 
    [cp client to /usr/bin/]
    Password:
    [get stat] [curl 127.0.0.1:9090/status/]
    {"Count":0,"KeySize":0,"ValueSize":0}
    
    [set cache] [client -s 127.0.0.1:9091 -c set -k testkey -v testvalue]
    testvalue
    
    [get cache] [client -s 127.0.0.1:9091 -c get -k testkey]
    testvalue
    
    [get stat] [curl 127.0.0.1:9090/status/]
    {"Count":1,"KeySize":7,"ValueSize":9}
    
    [delete cache] [client -s 127.0.0.1:9091 -c del -k testkey]
    
    
    [get stat] [curl 127.0.0.1:9090/status/]
    {"Count":0,"KeySize":0,"ValueSize":0}
    
    [remove client from /usr/bin/]

    ```

3. 性能测试

    测试机器需要安装 [Redis](https://redis.io)，从测试结果看：
    
    - 写：Redis 性能大约在自建的 Http 缓存服务的 1.4 倍左右。
    - 读：Redis 性能大约在自建的 Http 缓存服务的 1.7 倍左右。

    ```shell
    # 编译
    ➜  github.com/jormin/go-cache git:(master) cd cache-benchmark
    ➜  server git:(master) ✗ go build
    # 运行测试脚本
    ➜  github.com/jormin/go-cache git:(master) cd chapter2/test
    ➜  test git:(master) ✗ ./benchmark.sh
    [cp cache-benchmark to /usr/bin/]
    [cache-set] [cache-benchmark -type tcp -h 127.0.0.1:9091 -n 10000 -r 10000 -t set]
    type is tcp
    server is 127.0.0.1:9091
    total 10000 requests
    data size is 1000
    we have 1 connections
    operation is set
    keyspacelen is 10000
    pipeline length is 1
    0 records get
    0 records miss
    10000 records set
    0.387958 seconds total
    100% requests < 1 ms
    36 usec average for each request
    throughput is 25.775965 MB/s
    rps is 25775.964733
    
    [cache-get] [cache-benchmark -type tcp -h 127.0.0.1:9091 -n 10000 -r 10000 -t get]
    type is tcp
    server is 127.0.0.1:9091
    total 10000 requests
    data size is 1000
    we have 1 connections
    operation is get
    keyspacelen is 10000
    pipeline length is 1
    6286 records get
    3714 records miss
    0 records set
    0.453094 seconds total
    100% requests < 1 ms
    37 usec average for each request
    throughput is 13.873507 MB/s
    rps is 22070.485233
    
    [redis-set,get] [redis-benchmark -c 1 -n 10000 -d 1000 -t set,get -r 10000]
    ====== SET ======
      10000 requests completed in 0.27 seconds
      1 parallel clients
      1000 bytes payload
      keep alive: 1
      host configuration "save": 3600 1 300 100 60 10000
      host configuration "appendonly": no
      multi-thread: no
    
    100.00% <= 0.1 milliseconds
    36630.04 requests per second
    
    ====== GET ======
      10000 requests completed in 0.27 seconds
      1 parallel clients
      1000 bytes payload
      keep alive: 1
      host configuration "save": 3600 1 300 100 60 10000
      host configuration "appendonly": no
      multi-thread: no
    
    99.90% <= 0.1 milliseconds
    100.00% <= 0.4 milliseconds
    37037.04 requests per second
    
    
    
    [remove cache-benchmark from /usr/bin/]

    ```