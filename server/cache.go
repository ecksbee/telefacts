package server

import (
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

var (
	once     sync.Once
	appcache *gocache.Cache
)

func NewCache() *gocache.Cache {
	once.Do(func() {
		appcache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
	})
	return appcache
}
