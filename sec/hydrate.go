package sec

import (
	"fmt"

	"ecksbee.com/telefacts/hydratables"
	"ecksbee.com/telefacts/serializables"
	gocache "github.com/patrickmn/go-cache"
)

func Hydrate(workingDir string, folder *serializables.Folder) (*hydratables.Hydratable, error) {
	ret, err := hydratables.Hydrate(folder)
	if err != nil {
		return nil, err
	}
	go func() {
		lock.Lock()
		defer lock.Unlock()
		appcache.Set(workingDir, *ret, gocache.DefaultExpiration)
	}()
	return ret, nil
}

func Hydratable(workingDir string) (*hydratables.Hydratable, error) {
	if appcache == nil {
		return nil, fmt.Errorf("No accessible cache")
	}
	lock.RLock()
	defer lock.RUnlock()
	if x, found := appcache.Get(workingDir); found {
		ret := x.(hydratables.Hydratable)
		return &ret, nil
	}
	folder, err := Folder(workingDir)
	if err != nil {
		return nil, fmt.Errorf("Failed to discover folder")
	}
	return Hydrate(workingDir, folder)
}
