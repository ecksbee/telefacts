package telefacts_test

import (
	"sync"
	"time"
)

const (
	SEC_INTERVAL = time.Second * 2
)

var (
	secMutex sync.Mutex
)
