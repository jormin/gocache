#!/bin/bash

# 查询缓存状态
echo "[get stat] [curl 127.0.0.1:9090/status/]"
curl 127.0.0.1:9090/status/ -w '\n\n'

# 设置缓存
echo "[set cache] [curl -v 127.0.0.1:9090/cache/testkey -XPUT -dtestvalue]"
curl -v 127.0.0.1:9090/cache/testkey -XPUT -dtestvalue -w '\n\n'

# 读取缓存
echo "[get cache] [curl 127.0.0.1:9090/cache/testkey]"
curl 127.0.0.1:9090/cache/testkey -w '\n\n'

# 查询缓存状态
echo "[get stat] [curl 127.0.0.1:9090/status/]"
curl 127.0.0.1:9090/status/ -w '\n\n'

# 删除缓存
echo "[delete cache] [curl 127.0.0.1:9090/cache/testkey -XDELETE]"
curl 127.0.0.1:9090/cache/testkey -XDELETE -w '\n\n'

# 查询缓存状态
echo "[get stat] [curl 127.0.0.1:9090/status/]"
curl 127.0.0.1:9090/status/ -w '\n\n'