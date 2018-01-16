package base

import (
	"github.com/muesli/cache2go"
	"time"
	"sync"
)

type Cache struct {
	table *cache2go.CacheTable
}

var cache *Cache
var cacheOnce sync.Once

func Default() *Cache  {
	cacheOnce.Do(func() {
		cache = new(Cache)
		cache.table = cache2go.Cache("default")
	})
	return cache
}

func (this *Cache) Add(key string, val interface{}, t time.Duration) {
	this.table.Add(key, t, val)
}


func (this *Cache) Remove(key string) interface{} {
	val, err := this.table.Delete(key)
	if err != nil {
		return nil
	}
	return val.Data()
}

func (this *Cache) Has(key string) bool {
	return this.table.Exists(key)
}

func (this *Cache) Get(key string) interface{} {
	val, err := this.table.Value(key)
	if err != nil {
		return nil
	}
	return val.Data()
}
