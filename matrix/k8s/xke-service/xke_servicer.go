package xke_service

import "dtstack.com/dtstack/easymatrix/matrix/k8s/xke-service/driver"

type XkeServicer interface {
	Create(name, config string, id int) error
	Deploy(name, yaml string) error
	DeployWithF(name, file string) error
}

type xkeService struct {
}

func NewXkeService() (XkeServicer, error) {
	newService := &xkeService{}
	return newService, nil
}

func (this *xkeService) Create(clusterName, config string, clusterId int) error {
	return driver.RkeCreate(clusterName, config, clusterId)
}

func (this *xkeService) Deploy(clusterName, yaml string) error {
	return driver.DeployWithKubeCtl(clusterName, yaml)
}

func (this *xkeService) DeployWithF(clusterName, file string) error {
	return driver.DeployWithKubeCtlWithFile(clusterName, file)
}
