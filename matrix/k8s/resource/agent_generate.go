package resource

import (
	"bytes"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s/view"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/cluster"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/constant"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource/secret"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	modelkube "dtstack.com/dtstack/easymatrix/matrix/model/kube"
	"encoding/base64"
	"text/template"
)
var importClusterNsTemplateFiles = constant.TPL_NS_RESOURCE

func AgentGenerate(req *view.AgentGenerateReq, clusterId int) (*view.AgentGenerateRsp,error){
	tpl, err := cluster.ReadAndParseTemplate(importClusterNsTemplateFiles.FileName)
	if err != nil{
		return nil,err
	}
	registry := ""
	dockerConfigJson := ""
	tbsc,err := modelkube.DeployClusterImageStore.GetById(req.RegistryId)
	clusterImageStore, err := modelkube.DeployClusterImageStore.GetByClusterId(clusterId)
	if err != nil{
		return nil,err
	}
	if tbsc != nil{
		registry = tbsc.Address+"/"
		se,err := secret.GetDockerConfigJson(clusterImageStore,req.Namespace,tbsc.Alias)
		if err != nil{
			return nil,err
		}
		dockerConfigJson = base64.StdEncoding.EncodeToString(se.Data[".dockerconfigjson"])
	}
	bts,err := generateTplFile(tpl,dockerConfigJson,req.Namespace,registry)
	if err != nil{
		return nil,err
	}
	return &view.AgentGenerateRsp{
		Yaml: string(bts),
	},nil
}

func generateTplFile(tpl *template.Template, secret,namespace,registry string) ([]byte,error){
	var buf bytes.Buffer
	err := tpl.Execute(&buf,map[string]string{
		"NAME_SPACE": namespace,
		"REGISTRY": registry,
		"IMAGE_SECRET":secret,
	})
	if err != nil{
		log.Errorf("[import_cluster_ns]: template execute error : %v",err)
		return nil,err
	}
	return buf.Bytes(),nil
}
