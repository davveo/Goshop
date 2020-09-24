package cache

import (
	"time"

	"github.com/allegro/bigcache"
)

var (
	config = bigcache.Config{
		Shards:             16,              // 存储的条目数量，值必须是2的幂
		LifeWindow:         2 * time.Second, // 超时后条目被删除
		MaxEntriesInWindow: 0,               // 在 Life Window 中的最大数量
		MaxEntrySize:       0,               // 条目最大尺寸，以字节为单位
		HardMaxCacheSize:   0,               // 设置缓存最大值，以MB为单位，超过了不在分配内存。0表示无限制分配
	}
	Cache, _ = bigcache.NewBigCache(config)
)

func GetCache() *bigcache.BigCache {
	cache, _ := bigcache.NewBigCache(config)
	return cache
}
