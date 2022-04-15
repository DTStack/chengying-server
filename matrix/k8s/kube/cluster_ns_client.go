package kube

import (
	"fmt"
	"sync"
)

var ClusterNsClientCache  = &clusterNsClientCache{}

type NsClientCacheEnhance struct {
	kubeClient *KubeClientCache
	restClient *RestClientCache
}

func (c *NsClientCacheEnhance) GetClientCache(typ ImportType) (ClientCache,error){
	switch typ {
	case IMPORT_KUBECONFIG:
		return c.kubeClient,nil
	case IMPORT_AGENT:
		return c.restClient,nil
	default:
		return nil,fmt.Errorf("[cluster_ns_client]: unknow import type %v",typ)
	}
}

func (c *NsClientCacheEnhance) DeleteNsClient(namespace string){
	c.kubeClient.DeleteClient(namespace)
	c.restClient.DeleteClient(namespace)
}

func (c *NsClientCacheEnhance) PutNsClient(cache ClientCache){
	kubeC, ok := cache.(*KubeClientCache)
	if ok{
		c.kubeClient = kubeC
	}else {
		restC := cache.(*RestClientCache)
		c.restClient = restC
	}
}

type clusterNsClientCache struct {
	clusterNsClient map[string]*NsClientCacheEnhance
	mu sync.RWMutex
}

func (c *clusterNsClientCache)GetClusterNsClient(clusterName string) *NsClientCacheEnhance{
	c.mu.RLock()
	nsClientCache,konw := c.clusterNsClient[clusterName]
	c.mu.RUnlock()
	if konw{
		return nsClientCache
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.clusterNsClient == nil{
		c.clusterNsClient = make(map[string]*NsClientCacheEnhance)
	}
	c.clusterNsClient[clusterName] = &NsClientCacheEnhance{
		kubeClient: &KubeClientCache{},
		restClient: &RestClientCache{},
	}
	return c.clusterNsClient[clusterName]
}


