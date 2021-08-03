package hydratables

import (
	"fmt"
	"sync"

	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

var (
	lock        sync.RWMutex
	globalCache *gocache.Cache
)

func InjectCache(c *gocache.Cache) {
	globalCache = c
}

func HydrateGlobalSchema(urlStr string) (*Schema, error) {
	if globalCache == nil {
		return nil, fmt.Errorf("no accessible cache")
	}
	lock.RLock()
	if x, found := globalCache.Get(urlStr); found {
		ret := x.(Schema)
		lock.RUnlock()
		return &ret, nil
	}
	lock.RUnlock()
	file, err := serializables.DiscoverGlobalSchema(urlStr)
	if err != nil {
		return nil, err
	}
	schema, err := HydrateSchema(file, urlStr)
	if err != nil {
		return nil, err
	}
	go func() {
		lock.Lock()
		defer lock.Unlock()
		globalCache.Set(urlStr, *schema, gocache.DefaultExpiration)
	}()
	return schema, err
}

func HydrateFundamentalSchema() (*Schema, error) {
	return HydrateGlobalSchema(attr.LRR)
}

func HydrateEntityNames() (map[string]map[string]string, error) {
	key := "names.json"
	if globalCache == nil {
		return nil, fmt.Errorf("no accessible cache")
	}
	lock.RLock()
	if x, found := globalCache.Get(key); found {
		ret := x.(map[string]map[string]string)
		lock.RUnlock()
		return ret, nil
	}
	lock.RUnlock()
	names, err := serializables.DiscoverEntityNames()
	if err != nil {
		return names, err
	}
	go func() {
		lock.Lock()
		defer lock.Unlock()
		globalCache.Set(key, names, gocache.DefaultExpiration)
	}()
	return names, err
}
