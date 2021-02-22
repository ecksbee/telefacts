package server

import (
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

var (
	sOnce    sync.Once
	sCache   *gocache.Cache
	hOnce    sync.Once
	hCache   *gocache.Cache
	secOnce  sync.Once
	secCache *gocache.Cache
)

func NewSCache() *gocache.Cache {
	sOnce.Do(func() {
		sCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	})
	return sCache
}

func NewHCache() *gocache.Cache {
	hOnce.Do(func() {
		hCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	})
	return hCache
}

func NewSECCache() *gocache.Cache {
	secOnce.Do(func() {
		secCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	})
	return secCache
}
