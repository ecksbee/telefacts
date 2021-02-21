package hydratables

import (
	"fmt"
	"sync"

	"ecksbee.com/telefacts/serializables"
	gocache "github.com/patrickmn/go-cache"
)

var (
	lock     sync.RWMutex
	appcache *gocache.Cache
)

func InjectCache(c *gocache.Cache) {
	appcache = c
}

func HydrateGlobalSchema(urlStr string) (*Schema, error) {
	if appcache == nil {
		return nil, fmt.Errorf("No accessible cache")
	}
	lock.RLock()
	if x, found := appcache.Get(urlStr); found {
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
		appcache.Set(urlStr, *schema, gocache.DefaultExpiration)
	}()
	return schema, err
}
