package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	CreateSessionId = createSessionId()
	sessionIdCache  = cache.New(10*time.Minute, 5*time.Second)
	LoginLockStatus = cache.New(10*time.Minute, 5*time.Second)
)
