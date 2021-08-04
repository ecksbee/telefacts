package cache

import (
	"fmt"
	"sync"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/renderables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

var (
	lock     sync.RWMutex
	once     sync.Once
	appCache *gocache.Cache
)

func NewCache() *gocache.Cache {
	once.Do(func() {
		appCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	})
	return appCache
}

func Marshal(id string, hash string) ([]byte, error) {
	if appCache == nil {
		return nil, fmt.Errorf("no accessible cache")
	}
	h, err := hydratable(id)
	if err != nil {
		return nil, err
	}
	lock.RLock()
	if x, found := appCache.Get(id + "/" + hash); found {
		ret := x.([]byte)
		lock.RUnlock()
		return ret, nil
	}
	lock.RUnlock()
	byteArr, err := renderables.MarshalRenderable(hash, h)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(id+"/", byteArr, gocache.DefaultExpiration)
	}()
	return byteArr, err
}

func MarshalCatalog(id string) ([]byte, error) {
	if appCache == nil {
		return nil, fmt.Errorf("no accessible cache")
	}
	h, err := hydratable(id)
	if err != nil {
		return nil, err
	}
	lock.RLock()
	if x, found := appCache.Get(id + "/"); found {
		ret := x.([]byte)
		lock.RUnlock()
		return ret, nil
	}
	lock.RUnlock()
	byteArr, err := renderables.MarshalCatalog(h)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(id+"/", byteArr, gocache.DefaultExpiration)
	}()
	return byteArr, err
}

func hydratable(id string) (*hydratables.Hydratable, error) {
	lock.RLock()
	if x, found := appCache.Get(id); found {
		ret := x.(*hydratables.Hydratable)
		lock.RUnlock()
		return ret, nil
	}
	lock.RUnlock()
	folder, err := serializables.Discover(id)
	if err != nil {
		return nil, fmt.Errorf("failed to discover folder, %v", err)
	}
	ret, err := hydratables.Hydrate(folder)
	go func() {
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(id, ret, gocache.DefaultExpiration)
	}()
	return ret, err
}
