package cache

import "sync"

type inMemoryCache struct {
	c map[string][]byte
	// 读写锁
	mutex sync.RWMutex
	Stat
}

// 设置缓存
func (i *inMemoryCache) Set(k string, v []byte) error {
	// 上锁
	i.mutex.Lock()
	// 函数执行完毕后解锁
	i.mutex.Unlock()
	// 判断是否存在
	tmp, exist := i.c[k]
	// 如果存在先删除
	if exist {
		// 调用 Stat 的删除方法
		i.del(k, tmp)
		// 删除
		delete(i.c, k)
	}
	// 设置
	i.c[k] = v
	// 调用 Stat 的添加方法
	i.add(k, v)
	return nil
}

// 读取缓存
func (i *inMemoryCache) Get(k string) ([]byte, error) {
	// 读锁
	i.mutex.RLock()
	// 函数执行完毕后解除读锁
	defer i.mutex.RUnlock()
	// 读取数据并返回
	return i.c[k], nil
}

// 删除
func (i *inMemoryCache) Delete(k string) error {
	// 上锁
	i.mutex.Lock()
	// 函数执行完毕后解锁
	defer i.mutex.Unlock()
	tmp, exist := i.c[k]
	if exist {
		// 调用 Stat 的删除方法
		i.del(k, tmp)
		// 删除
		delete(i.c, k)
	}
	return nil
}

// 获取状态
func (i *inMemoryCache) GetStat() Stat {
	return i.Stat
}

// 生成内存缓存
func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{
		make(map[string][]byte),
		sync.RWMutex{},
		Stat{},
	}
}
