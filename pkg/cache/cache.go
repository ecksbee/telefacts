package cache

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"path/filepath"
	"sync"

	"ecksbee.com/telefacts/pkg/hydratables"
	"ecksbee.com/telefacts/pkg/renderables"
	"ecksbee.com/telefacts/pkg/serializables"
	gocache "github.com/patrickmn/go-cache"
)

var (
	dry      bool
	lock     sync.RWMutex
	once     sync.Once
	appCache *gocache.Cache
)

func NewCache(runDry bool) *gocache.Cache {
	dry = runDry
	if dry {
		return appCache
	}
	once.Do(func() {
		appCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	})
	return appCache
}

func MarshalExpressable(id string, name string, contextref string) ([]byte, error) {
	h, err := hydratable(id)
	if err != nil {
		return nil, err
	}
	hash := fnv.New128a()
	hash.Write([]byte(id + "/facts/" + name + ":" + contextref))
	cachekey := hex.EncodeToString(hash.Sum([]byte{}))
	lock.RLock()
	if !dry {
		if x, found := appCache.Get(cachekey); found {
			ret := x.([]byte)
			lock.RUnlock()
			return ret, nil
		}
	}
	lock.RUnlock()
	byteArr, err := renderables.MarshalExpressable(name, contextref, h)
	go func() {
		if dry {
			return
		}
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(cachekey, byteArr, gocache.DefaultExpiration)
	}()
	return byteArr, err
}

func MarshalRenderable(id string, hash string) ([]byte, error) {
	lock.RLock()
	if !dry {
		if x, found := appCache.Get(id + "/" + hash); found {
			ret := x.([]byte)
			lock.RUnlock()
			return ret, nil
		}
	}
	lock.RUnlock()
	h, err := hydratable(id)
	if err != nil {
		return nil, err
	}
	ext := filepath.Ext(hash)
	if ext == ".xhtml" || ext == ".htm" {
		entryFileName := h.Folder.EntryFileName
		if entryFileName == hash {
			data := h.Folder.Document.Bytes
			if dry {
				return data, nil
			}
			go func() {
				lock.Lock()
				defer lock.Unlock()
				appCache.Set(id+"/"+hash, data, gocache.DefaultExpiration)
			}()
			return data, nil
		}
	}
	byteArr, err := renderables.MarshalRenderable(hash, h)
	go func() {
		if dry {
			return
		}
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(id+"/"+hash, byteArr, gocache.DefaultExpiration)
	}()
	return byteArr, err
}

func MarshalCatalog(id string) ([]byte, error) {
	h, err := hydratable(id)
	if err != nil {
		return nil, err
	}
	lock.RLock()
	if !dry {
		if x, found := appCache.Get(id + "/"); found {
			ret := x.([]byte)
			lock.RUnlock()
			return ret, nil
		}
	}
	lock.RUnlock()
	byteArr, err := renderables.MarshalCatalog(h)
	go func() {
		if dry {
			return
		}
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(id+"/", byteArr, gocache.DefaultExpiration)
	}()
	return byteArr, err
}

func hydratable(id string) (*hydratables.Hydratable, error) {
	lock.RLock()
	if !dry {
		if x, found := appCache.Get(id); found {
			ret := x.(*hydratables.Hydratable)
			lock.RUnlock()
			return ret, nil
		}
	}
	lock.RUnlock()
	folder, err := serializables.Discover(id)
	if err != nil {
		return nil, fmt.Errorf("failed to discover folder, %v", err)
	}
	ret, err := hydratables.Hydrate(folder)
	go func() {
		if dry {
			return
		}
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(id, ret, gocache.DefaultExpiration)
	}()
	return ret, err
}
