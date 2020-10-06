#!/bin/bash

# 将 tcp 客户端 cp 到 bin 目录
echo "[cp client to /usr/bin/]"
sudo cp ../../client/client /usr/bin/client

echo "[get stat] [curl 127.0.0.1:9090/status/]"
curl 127.0.0.1:9090/status/ -w '\n\n'

# 设置缓存
echo "[set cache] [client -s 127.0.0.1:9091 -c set -k testkey -v testvalue]"
client -s 127.0.0.1:9091 -c set -k testkey -v testvalue

# 读取缓存
echo ""
echo "[get cache] [client -s 127.0.0.1:9091 -c get -k testkey]"
client -s 127.0.0.1:9091 -c get -k testkey

# 查询缓存状态
echo ""
echo "[get stat] [curl 127.0.0.1:9090/status/]"
curl 127.0.0.1:9090/status/ -w '\n\n'

# 删除缓存
echo "[delete cache] [client -s 127.0.0.1:9091 -c del -k testkey]"
client -s 127.0.0.1:9091 -c del -k testkey

# 查询缓存状态
echo ""
echo "[get stat] [curl 127.0.0.1:9090/status/]"
curl 127.0.0.1:9090/status/ -w '\n\n'

# 删除 tcp 客户端
echo "[remove client from /usr/bin/]"
sudo rm -f /usr/bin/client