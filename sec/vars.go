package sec

import (
	"sync"

	gocache "github.com/patrickmn/go-cache"
)

var (
	lock     sync.RWMutex
	appcache *gocache.Cache
)

func InjectCache(c *gocache.Cache) {
	appcache = c
}
