package deploy

import (
	oldclient "dtstack.com/dtstack/easymatrix/addons/oldkube/pkg/client-go"
	"dtstack.com/dtstack/easymatrix/matrix/agent"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func ApplyDynamicResource(d *oldclient.DynamicData, clusterId int) error {
	log.Infof("k8s deploy apply dynamic resource ...")
	body, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("json marshal error:%v", err)
	}
	log.Infof("%v", string(body[:]))

	content := oldclient.ContentResponse{}
	clientParam := agent.ExecRestParams{
		Method:  "POST",
		Path:    "clientgo/dynamic/apply",
		Body:    body,
		Timeout: "5s",
	}
	cluster, _ := model.DeployClusterList.GetClusterInfoById(clusterId)
	sid, _ := model.DeployNodeList.GetDeployNodeSidByClusterIdAndMode(clusterId, cluster.Mode)
	resp, err := agent.AgentClient.ToExecRest(sid, &clientParam, "")
	if err != nil {
		return fmt.Errorf("ToExecRest dynamic apply err:%v", err)
	}
	decodeResp, err := base64.URLEncoding.DecodeString(resp)
	if err != nil {
		log.Errorf("client-go response decode err:%v", err)
	}
	_ = json.Unmarshal(decodeResp, &content)
	if content.Code != 0 {
		return fmt.Errorf(content.Msg)
	}
	return nil
}

func NewDynamic(data []byte, gvr *schema.GroupVersionResource, kind string) *oldclient.DynamicData {
	d := &oldclient.DynamicData{
		Data:     string(data),
		Group:    gvr.Group,
		Resource: gvr.Resource,
		Version:  gvr.Version,
	}
	return d
}

func DeleteDynamicResource(d *oldclient.DynamicData, clusterId int) error {
	log.Infof("k8s delete dynamic resource ...")
	body, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("json marshal error:%v", err)
	}
	log.Infof("%v", string(body[:]))

	content := oldclient.ContentResponse{}
	clientParam := agent.ExecRestParams{
		Method:  "POST",
		Path:    "clientgo/dynamic/delete",
		Body:    body,
		Timeout: "5s",
	}
	cluster, _ := model.DeployClusterList.GetClusterInfoById(clusterId)
	sid, _ := model.DeployNodeList.GetDeployNodeSidByClusterIdAndMode(clusterId, cluster.Mode)
	resp, err := agent.AgentClient.ToExecRest(sid, &clientParam, "")
	if err != nil {
		return fmt.Errorf("ToExecRest dynamic delete err:%v", err)
	}
	decodeResp, err := base64.URLEncoding.DecodeString(resp)
	if err != nil {
		log.Errorf("client-go response decode err:%v", err)
	}
	_ = json.Unmarshal(decodeResp, &content)
	if content.Code != 0 {
		return fmt.Errorf(content.Msg)
	}
	return nil
}

func GetDynamicResource(d *oldclient.DynamicData, clusterId int) (interface{}, error) {
	log.Infof("k8s get dynamic resource ...")
	body, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("json marshal error:%v", err)
	}
	log.Infof("%v", string(body[:]))

	content := oldclient.ContentResponse{}
	clientParam := agent.ExecRestParams{
		Method:  "POST",
		Path:    "clientgo/dynamic/get",
		Body:    body,
		Timeout: "5s",
	}
	cluster, _ := model.DeployClusterList.GetClusterInfoById(clusterId)
	sid, _ := model.DeployNodeList.GetDeployNodeSidByClusterIdAndMode(clusterId, cluster.Mode)
	resp, err := agent.AgentClient.ToExecRest(sid, &clientParam, "")
	if err != nil {
		return nil, fmt.Errorf("ToExecRest dynamic get err:%v", err)
	}
	decodeResp, err := base64.StdEncoding.DecodeString(resp)
	if err != nil {
		log.Errorf("client-go response decode err:%v", err)
	}
	_ = json.Unmarshal(decodeResp, &content)
	if content.Code != 0 {
		return nil, fmt.Errorf(content.Msg)
	}
	return content.Data, nil
}
