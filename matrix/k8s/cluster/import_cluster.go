package cluster

import (
	"dtstack.com/dtstack/easymatrix/matrix/asset"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/constant"
	"dtstack.com/dtstack/easymatrix/matrix/log"
)

var importClusterTemplateFiles = []constant.TemplateFile{
	constant.TPL_CLUSTER_RESOURCE,
	constant.TPL_CLUSTER_RESOURCE_V1BETA1,
}

type ImportClusterGenerator struct {

}

func (g *ImportClusterGenerator) Generate() (map[string][]byte,error){
	yamls := make(map[string][]byte,len(importClusterTemplateFiles))
	asset.ResetImportClusterTemplateWithLocalFile()
	for _,tplName := range importClusterTemplateFiles{
		bts,err := asset.Asset(tplName.FileName)
		if err != nil{
			log.Errorf("[import_cluster]: read cluster resource %s, error : %v",tplName.FileName,err)
			return nil ,err
		}
		yamls[tplName.FileName] = bts
	}
	return yamls,nil
}

func (g *ImportClusterGenerator) GetFileNames() []constant.TemplateFile{
	return importClusterTemplateFiles
}
