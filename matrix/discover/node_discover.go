package discover

import (
	"bytes"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"fmt"
	"io/ioutil"
	"sync"
	"text/template"

	"dtstack.com/dtstack/easymatrix/matrix/base"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"time"
)

var (
	nodeLock         = sync.Mutex{}
	NodeExporterPort int
)

const (
	NODE_SD_FILE = "/prometheus/node_sd_file.yml"

	NODE_SD_TPL = `{{range $_ := .}}
- targets: ['{{.Ip}}:%d']
  labels:
    sid: {{.Sid}}
    cluster_name: {{.ClusterName}}
    clusterId: {{.ClusterId}}
    type: {{.Type}}
{{end}}`
)

func FlushNodeDiscover() {
	nodeLock.Lock()
	defer nodeLock.Unlock()

	type ClusterInfo struct {
		Sid         string    `db:"id"`
		Ip          string    `db:"local_ip"`
		ClusterName string    `db:"cluster_name"`
		ClusterId   int       `db:"clusterId"`
		Type        string    `db:"type"`
		UpdateDate  base.Time `db:"last_update_date"`
	}
	clusterInfo := make([]ClusterInfo, 0)
	query := "SELECT h.id, deploy_cluster_list.name as cluster_name, clusterId, h.local_ip, type, h.last_update_date FROM deploy_cluster_host_rel " +
		"LEFT JOIN sidecar_list as h ON deploy_cluster_host_rel.sid = h.id " +
		"LEFT JOIN deploy_cluster_list ON deploy_cluster_list.id = deploy_cluster_host_rel.clusterId " +
		"WHERE deploy_cluster_host_rel.is_deleted=0 AND deploy_cluster_list.is_deleted=0"
	if err := model.USE_MYSQL_DB().Select(&clusterInfo, query); err != nil {
		log.Errorf("%v", err)
		return
	}
	freshClusterInfo := make([]ClusterInfo, 0)
	for _, info := range clusterInfo {
		if time.Now().Sub(time.Time(info.UpdateDate)) < 3*time.Minute {
			freshClusterInfo = append(freshClusterInfo, info)
		}
	}
	buf := &bytes.Buffer{}
	tpl := template.Must(template.New("node_discover").Option("missingkey=error").Parse(fmt.Sprintf(NODE_SD_TPL, NodeExporterPort)))
	if err := tpl.Execute(buf, freshClusterInfo); err != nil {
		log.Errorf("%v", err)
		return
	}

	if err := ioutil.WriteFile(NODE_SD_FILE, buf.Bytes(), 0755); err != nil {
		log.Errorf("%v", err)
		return
	}
}
