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

func Marshal(workingDir string, hash string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	return renderables.MarshalRenderable(hash, h)
}

func MarshalCatalog(workingDir string) ([]byte, error) {
	h, err := Hydratable(workingDir)
	if err != nil {
		return nil, err
	}
	filenames := []string{}
	return renderables.MarshalCatalog(h, filenames)
}

func Hydratable(workingDir string) (*hydratables.Hydratable, error) {
	if appCache == nil {
		return nil, fmt.Errorf("no accessible cache")
	}
	lock.RLock()
	defer lock.RUnlock()
	if x, found := appCache.Get(workingDir); found {
		ret := x.(hydratables.Hydratable)
		return &ret, nil
	}
	folder, err := serializables.Discover(workingDir)
	if err != nil {
		return nil, fmt.Errorf("failed to discover folder, %v", err)
	}
	return Hydrate(workingDir, folder)
}

func Hydrate(workingDir string, folder *serializables.Folder) (*hydratables.Hydratable, error) {
	ret, err := hydratables.Hydrate(folder)
	if err != nil {
		return nil, err
	}
	go func() {
		lock.Lock()
		defer lock.Unlock()
		appCache.Set(workingDir, *ret, gocache.DefaultExpiration)
	}()
	return ret, nil
}
