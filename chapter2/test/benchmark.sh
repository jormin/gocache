#!/bin/bash

# 将测试脚本 link 到 bin 目录
echo "[cp cache-benchmark to /usr/bin/]"
sudo cp ../../cache-benchmark/cache-benchmark /usr/bin/cache-benchmark

# cache 写
echo "[cache-set] [cache-benchmark -type http -n 10000 -r 10000 -t set]"
cache-benchmark -type http -n 10000 -r 10000 -t set

echo ""
# cache 读
echo "[cache-get] [cache-benchmark -type http -n 10000 -r 10000 -t get]"
cache-benchmark -type http -n 10000 -r 10000 -t get

echo ""
# redis 读写
echo "[redis-set,get] [redis-benchmark -c 1 -n 10000 -d 1000 -t set,get -r 10000]"
redis-benchmark -c 1 -n 10000 -d 1000 -t set,get -r 10000

echo ""
# 将测试脚本 link 到 bin 目录
echo "[remove cache-benchmark from /usr/bin/]"
sudo rm -f /usr/bin/cache-benchmark