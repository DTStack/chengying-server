package impl

import (
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
)

var clientCache *base.ClientCache

func init() {
	clientCache = &base.ClientCache{}
	//use incluseterconfig build clientcache
	clientCache.Connect("", "")
}
