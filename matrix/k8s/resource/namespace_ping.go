package resource

import (
	"context"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s/view"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/kube"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource/configmap"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource/deployment"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource/mole"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource/resourcequota"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource/secret"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource/service"
	"strings"
)

func NamespacePing(ctx context.Context,clusterid string,vo *view.NamespacePingReq) error{
	cache,err := kube.ClusterNsClientCache.GetClusterNsClient(clusterid).GetClientCache(kube.IMPORT_AGENT)
	if err != nil{
		return err
	}
	ip := vo.Ip
	port := vo.Port
	if !strings.HasPrefix(ip,"http"){
		ip = "http://"+ip
	}
	host := ip + ":" +port
	ns := vo.Namespace
	err = cache.Connect(host,ns)
	if err != nil{
		return err
	}
	client := cache.GetClient(ns)
	return ping(client,ns)
}

func ping(client kube.Client, ns string) error{
	var err error
	if err = service.Ping(client,ns); err != nil{
		return err
	}
	if err = secret.Ping(client,ns); err != nil{
		return err
	}
	if err = resourcequota.Ping(client,ns); err != nil{
		return err
	}
	if err = mole.Ping(client,ns); err != nil{
		return err
	}
	if err = deployment.Ping(client,ns); err != nil{
		return err
	}
	if err = configmap.Ping(client,ns); err != nil{
		return err
	}
	return nil
}
