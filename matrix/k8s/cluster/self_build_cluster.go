package cluster

import (
	"bytes"
	"dtstack.com/dtstack/easymatrix/matrix/asset"
	"dtstack.com/dtstack/easymatrix/matrix/host"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/constant"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"text/template"
)

var selfBuildTemplateFiles = []constant.TemplateFile{
	constant.TPL_SELF_BUILD,
}
type SelfBuilClusterGenerator struct {
	ClusterId string
	CallBackIp string
	Aid int
}

func (g *SelfBuilClusterGenerator) GetFileNames() []constant.TemplateFile{
	return selfBuildTemplateFiles
}

func (g *SelfBuilClusterGenerator) Generate() (map[string][]byte,error){
	yamls := make(map[string][]byte,len(selfBuildTemplateFiles))
	asset.ResetSelfBuildTemplateWithLocalFile()
	for _, tplName := range selfBuildTemplateFiles{
		tpl,err := ReadAndParseTemplate(tplName.FileName)
		if err != nil{
			return nil,err
		}
		bts,err := g.execute(tpl)
		if err != nil{
			return nil, err
		}
		yamls[tplName.FileName] = bts
	}
	return yamls,nil
}

func (g *SelfBuilClusterGenerator) execute(tpl *template.Template) ([]byte,error){
	callback, err := host.AgentInstall.GetAgentCallBack(g.Aid,g.CallBackIp)
	if err != nil{
		return nil,err
	}
	callback = callback+"&Mode=0&Deploy=daemonset&ClusterId="+g.ClusterId
	var buf bytes.Buffer
	err = tpl.Execute(&buf,map[string]string{
		"SERVER": host.AgentInstall.AgentHost,
		"CALLBACK": callback,
	})
	if err != nil{
		log.Errorf("[self_build_cluster]: template execute error : %v",err)
		return nil ,err
	}
	return buf.Bytes(),nil
}

