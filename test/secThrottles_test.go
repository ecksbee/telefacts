package telefacts_test

import (
	"sync"
	"time"
)

const (
	SEC_INTERVAL = time.Second * 5
)

var (
	secMutex sync.Mutex
)
