package cache

import (
	"sync"
	"time"
)

var LoginErrorCount = struct {
	Data map[string]int
	Lock sync.RWMutex
}{Data: make(map[string]int)}

func SetLoginCache(userName string, errCount int) {
	cleanTime := SysConfig.PlatFormSecurity.AccountLoginLockTime
	LoginLockStatus.Set(userName, errCount, time.Minute*time.Duration(cleanTime))
}
