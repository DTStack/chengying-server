package resource

import (
	"dtstack.com/dtstack/easymatrix/matrix/k8s/constant"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	modelkube "dtstack.com/dtstack/easymatrix/matrix/model/kube"
	"strconv"
)
type NamespaceListRsp struct {
	ClusterName		string 	`json:"cluster_name"`
	ClusterId		int 	`json:"cluster_id"`
	ClusterType		string 	`json:"cluster_type"`
	Namespaces		[]Namespace `json:"namespaces"`
}

type Namespace struct {
	Name 	string 	`json:"namespace"`
}
func NamespaceList(info *model.ClusterInfo) (*NamespaceListRsp,error) {
	nstbscs,err := modelkube.DeployNamespaceList.Select(strconv.Itoa(info.Id),constant.NAMESPACE_VALID,"","","")
	if err != nil{
		return nil,err
	}
	if nstbscs == nil {
		return nil,nil
	}
	nsList := []Namespace{}
	for _, tbsc := range nstbscs{
		ns := Namespace{
			Name: tbsc.Namespace,
		}
		nsList = append(nsList,ns)
	}
	return &NamespaceListRsp{
		ClusterId: info.Id,
		ClusterType: info.Type,
		ClusterName: info.Name,
		Namespaces: nsList,
	},nil
}
