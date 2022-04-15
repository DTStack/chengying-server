package cluster

import (
	"dtstack.com/dtstack/easymatrix/matrix/k8s/constant"
	modelkube "dtstack.com/dtstack/easymatrix/matrix/model/kube"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	info := &GeneratorInfo{
		Type:        constant.TYPE_SELF_BUILD,
		ClusterInfo: &modelkube.ClusterInfo{
			Id:            135,
			Name:          "em_mao_self",
		},
	}
	bts,_:=GetTemplateFile(info,false)
	fmt.Println(string(bts))
}

func Test2(t *testing.T) {

}
