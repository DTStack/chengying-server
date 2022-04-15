package kube

import "dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"

type KubeClientCache struct {
	c *base.ClientCache
}

func (c *KubeClientCache) GetClient(workspace string) Client{
	if c.c == nil{
		c.c = &base.ClientCache{}
	}
	return c.c.GetClient(workspace)
}

func (c *KubeClientCache) Connect(kubeconfig,workspace string) error{
	if c.c == nil{
		c.c = &base.ClientCache{}
	}
	return c.c.Connect(kubeconfig,workspace)
}

func (c *KubeClientCache) DeleteClient(workspace string){
	if c.c == nil{
		c.c = &base.ClientCache{}
	}
	c.c.DeleteClient(workspace)
}

func (c KubeClientCache) Copy() ClientCache{
	return &c
}
