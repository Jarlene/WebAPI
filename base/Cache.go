package base

import (
	"github.com/muesli/cache2go"
	"time"
)

type Cache struct {
	table *cache2go.CacheTable
}

var cache  = Cache{table:cache2go.Cache("default")}

func Default() *Cache  {
	return &cache
}


var SCache = Default()

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
