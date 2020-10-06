package cache

import "log"

// 创建并返回 Cache 接口
func New(typ string) Cache {
	var c Cache
	switch typ {
	case "inmemory":
		c = newInMemoryCache()
	default:
		panic("unknown cache type " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
