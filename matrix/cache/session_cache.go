package cache

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"time"
)

var (
	//	/api/v2/product/xxx/xxx/history
	//	/api/v2/cluster/restartServices
	//	/api/v2/product/anomalyService
	//	/api/v2/instance/product/xxx/service/xxx
	//	/api/v2/instance/product/xxx/service/xxx/healthCheck
	//	/api/v2/cluster/restartServices
	//	/api/v2/cluster/hosts/hosts
	//	/api/v2/cluster/hostgroups
	//	/api/v2/cluster/orderList
	//	/api/v2/cluster/orderDetail
	noFlushApiReg = []*regexp.Regexp{
		regexp.MustCompile("/api/v2/product/.*/.*/history"),
		regexp.MustCompile("/api/v2/cluster/restartServices"),
		regexp.MustCompile("/api/v2/product/anomalyService"),
		regexp.MustCompile("/api/v2/instance/product/.*/service/.*"),
		regexp.MustCompile("/api/v2/instance/product/.*/service/.*/healthCheck"),
		regexp.MustCompile("/api/v2/cluster/restartServices"),
		regexp.MustCompile("/api/v2/cluster/hosts/hosts"),
		regexp.MustCompile("/api/v2/cluster/hostgroups"),
		regexp.MustCompile("/api/v2/cluster/orderList"),
		regexp.MustCompile("/api/v2/cluster/orderDetail"),
	}
)

func SetSessionCache(sessionStr string) {
	cleanTime := SysConfig.PlatFormSecurity.AccountLogoutSleepTime
	sessionIdCache.Set(sessionStr, true, time.Minute*time.Duration(cleanTime))
}

func ValidationSessionId(path, hashId string) bool {
	_, ok := sessionIdCache.Get(hashId)
	if !ok {
		return false
	}
	for _, api := range noFlushApiReg {
		if api.MatchString(path) {
			return true
		}
	}
	SetSessionCache(hashId)
	return true
}

func createSessionId() (f func() string) {
	i := 0
	return func() string {
		w := md5.New()
		io.WriteString(w, strconv.Itoa(i))
		i++
		if i == math.MaxInt32 {
			i = 0
		}
		sessionStr := fmt.Sprintf("%x", w.Sum(nil))
		return sessionStr
	}
}
