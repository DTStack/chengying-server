package impl

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/base"
	"dtstack.com/dtstack/easymatrix/matrix/grafana"
	"dtstack.com/dtstack/easymatrix/matrix/harole"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"dtstack.com/dtstack/easymatrix/matrix/util"
	"encoding/json"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/kataras/iris/context"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
	"io/ioutil"
	"net/url"

	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var gofpdfDir string

func init() {
	gofpdfDir = "./"
}

func FontDir() string {
	return filepath.Join(gofpdfDir, "font")
}

func FontFile(fileStr string) string {
	return filepath.Join(FontDir(), fileStr)
}

func PdfDir(id int) string {
	return filepath.Join(base.WebRoot, "pdf", strconv.Itoa(id))
}

func PdfFile(id int, fileStr string) string {
	return filepath.Join(PdfDir(id), fileStr)
}

func ImageDir(id int) string {
	return filepath.Join(base.WebRoot, "img", strconv.Itoa(id))
}

func ImageFile(id int, fileStr string) string {
	return filepath.Join(ImageDir(id), fileStr)
}

func Filename(id int, baseStr string) string {
	return PdfFile(id, baseStr+".pdf")
}

type resultLists struct {
	ServiceName string `json:"service_name"`
	Ip          string `json:"ip"`
	Status      string `json:"status"`
	HARole      string `json:"ha_role"`
}

func GetServiceStatus(ctx context.Context) apibase.Result {
	// 获取当前所在父产品名称
	parentProduct, err := GetCurrentParentProduct(ctx)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	// 获取当前集群id
	clusterId, err := GetCurrentClusterId(ctx)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	productList, err := model.DeployProductList.GetDeploySonProductName(parentProduct, clusterId)
	if err != nil {
		log.Errorf("get son product name error: %v", err)
		return err
	}

	result := map[string][]resultLists{}
	for _, productName := range productList {
		instanceList, err := model.DeployInstanceList.FindByProductNameAndClusterId(productName, clusterId)
		if err != nil {
			log.Errorf("%v", err)
		}
		var resultInfoList []resultLists
		for _, instance := range instanceList {
			resultInfo := resultLists{
				ServiceName: instance.ServiceName,
				Ip:          instance.Ip,
			}
			status := model.INSTANCE_NORMAL
			if instance.Status != model.INSTANCE_STATUS_RUNNING {
				status = model.INSTANCE_ABNORMAL
			} else if instance.HealthState != model.INSTANCE_HEALTH_OK && instance.HealthState != model.INSTANCE_HEALTH_NOTSET {
				status = model.INSTANCE_ABNORMAL
			}
			resultInfo.Status = status
			roleData := harole.RoleData(instance.Pid, instance.ServiceName)
			if roleData != nil {
				haRole, ok := roleData[instance.AgentId]
				if !ok {
					haRole = "-"
				}
				resultInfo.HARole = strings.Replace(haRole, "\r", "", 1)
			}
			resultInfoList = append(resultInfoList, resultInfo)
		}
		result[productName] = resultInfoList
	}

	return result

}

type alertInfo struct {
	AlertName      string `json:"alert_name"`
	State          string `json:"state"`
	DashboardName  string `json:"dashboard_name"`
	DashboardTitle string `json:"dashboard_title"`
	Time           string `json:"time"`
}

func GetAlertHistory(ctx context.Context) apibase.Result {
	paramErrs := apibase.NewApiParameterErrors()
	from, err := ctx.URLParamInt("from")
	if err != nil {
		paramErrs.AppendError("$", "param from is empty")
	}
	to, err := ctx.URLParamInt("to")
	if err != nil {
		paramErrs.AppendError("$", "param to is empty")
	}
	paramErrs.CheckAndThrowApiParameterErrors()
	alertList := FormatAlertList(from, to)
	return map[string]interface{}{
		"count": len(alertList),
		"data":  alertList,
	}
}

func FormatAlertList(from, to int) []alertInfo {
	params := map[string]string{}
	params["from"] = strconv.Itoa(from)
	params["to"] = strconv.Itoa(to)
	params["type"] = "alert"
	params["limit"] = "1000"
	err, resp := grafana.GetAnnotations(params)
	if err != nil {
		log.Errorf("get annotations from grafana error: %v", err)
		return make([]alertInfo, 0)
	}

	var alertList []alertInfo

	for _, info := range resp {
		if info.NewState != "ok" && info.NewState != "paused" && info.NewState != "pending" {
			param := map[string]string{
				"dashboardId": strconv.Itoa(info.DashboardId),
				"panelId":     strconv.Itoa(info.PanelId),
			}
			err, alertRule := grafana.GrafanaAlertsSearch(param)
			if err != nil || len(alertRule) == 0 {
				log.Errorf("get alert rule: %v error: %v", info.AlertId, err)
				continue
			}
			panelTitle, dashboardName := RetrievePanelTitle(alertRule[0].DashboardUid, alertRule[0].PanelId)
			alert := alertInfo{
				AlertName:      info.AlertName,
				State:          info.NewState,
				DashboardName:  dashboardName,
				DashboardTitle: panelTitle,
				Time:           time.Unix(info.Time/1000, 0).Format("2006-01-02 15:04:05"),
			}
			alertList = append(alertList, alert)
		}
	}
	return alertList
}

type ReportHostInfo struct {
	Ip         string `json:"ip"`
	Cpu        string `json:"cpu"`
	Mem        string `json:"mem"`
	SystemDisk string `json:"system_disk"`
	DataDisk   string `json:"data_disk"`
}

func GetHostStatus(ctx context.Context) apibase.Result {
	// 获取当前所在父产品名称
	parentProduct, err := GetCurrentParentProduct(ctx)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	// 获取当前集群Id
	clusterId, err := GetCurrentClusterId(ctx)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}

	// 获取当前集群下所有接入的主机
	query := "select deploy_host.ip from deploy_cluster_host_rel " +
		"left join deploy_host on deploy_cluster_host_rel.sid=deploy_host.sid " +
		"left join deploy_instance_list on deploy_host.sid=deploy_instance_list.sid " +
		"left join deploy_product_list on deploy_instance_list.pid=deploy_product_list.id " +
		"left join sidecar_list on sidecar_list.id=deploy_host.sid where deploy_host.sid!='' " +
		"and deploy_host.isDeleted=0 and deploy_product_list.parent_product_name=? and deploy_cluster_host_rel.clusterId=? group by deploy_host.sid"
	var ipList []string
	if err := model.USE_MYSQL_DB().Select(&ipList, query, parentProduct, clusterId); err != nil {
		log.Errorf("get hosts error: %v", err)
		return err
	}

	hostStatusMap := map[string]ReportHostInfo{}
	// 初始化主机正常状态
	for _, ip := range ipList {
		hostStatusMap[ip] = ReportHostInfo{
			Ip:         ip,
			Cpu:        model.INSTANCE_NORMAL,
			Mem:        model.INSTANCE_NORMAL,
			SystemDisk: model.INSTANCE_NORMAL,
			DataDisk:   model.INSTANCE_NORMAL,
		}
	}

	// 获取Host Overview仪表盘信息
	err, dashboardResp := grafana.GetDashboardByUid("Ne_roaViz")
	if err != nil {
		log.Errorf("get host overview dashboard error: %v", err)
	}
	// 获取cpu告警信息
	err, cpuAlerts := grafana.AlertRuleTest(&grafana.AlertRuleTestParam{
		Dashboard: dashboardResp.Dashboard,
		PanelId:   38,
	})
	if err != nil {
		log.Errorf("test cpu alerts error: %v", err)
	}
	reg := regexp.MustCompile(`\w+-(?P<instance>(\d+\.)+\d+):\d+\s?(?P<mnt>/\w*)?`)
	if cpuAlerts.State != "ok" {
		for _, match := range cpuAlerts.Matches {
			metric := reg.FindStringSubmatch(match["metric"].(string))
			if metric != nil {
				if host, ok := hostStatusMap[metric[1]]; ok {
					host.Cpu = model.INSTANCE_ABNORMAL
					hostStatusMap[metric[1]] = host
				}
			}
		}
	}

	// 获取内存告警信息
	err, memoryAlerts := grafana.AlertRuleTest(&grafana.AlertRuleTestParam{
		Dashboard: dashboardResp.Dashboard,
		PanelId:   50,
	})
	if err != nil {
		log.Errorf("test memory alerts error: %v", err)
	}
	if memoryAlerts.State != "ok" {
		for _, match := range memoryAlerts.Matches {
			metric := reg.FindStringSubmatch(match["metric"].(string))
			if metric != nil {
				if host, ok := hostStatusMap[metric[1]]; ok {
					host.Mem = model.INSTANCE_ABNORMAL
					hostStatusMap[metric[1]] = host
				}
			}
		}

	}

	// 获取磁盘告警信息
	err, diskAlerts := grafana.AlertRuleTest(&grafana.AlertRuleTestParam{
		Dashboard: dashboardResp.Dashboard,
		PanelId:   44,
	})
	if err != nil {
		log.Errorf("test disk alerts error: %v", err)
	}
	if diskAlerts.State != "ok" {
		for _, match := range diskAlerts.Matches {
			metric := reg.FindStringSubmatch(match["metric"].(string))
			if metric != nil {
				if host, ok := hostStatusMap[metric[1]]; ok {
					if metric[3] == "/" {
						host.SystemDisk = model.INSTANCE_ABNORMAL
					}
					if metric[3] == "/data" {
						host.DataDisk = model.INSTANCE_ABNORMAL
					}
					hostStatusMap[metric[1]] = host
				}
			}
		}
	}

	var resultList = make([]ReportHostInfo, 0)
	for _, v := range hostStatusMap {
		resultList = append(resultList, v)
	}

	return resultList
}

func GetGraphConfig(ctx context.Context) apibase.Result {
	clusterId, err := GetCurrentClusterId(ctx)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	configList, err := model.InspectReportTemplate.GetTemplateConfig(clusterId)
	if err != nil {
		log.Errorf("%v", err)
	}
	return configList
}

type ChartsInfo struct {
	X []float64                `json:"x"`
	Y []map[string]interface{} `json:"y"`
}

type targetInfo struct {
	Expr         string `json:"expr"`
	LegendFormat string `json:"legend_format"`
}

func GetGraphData(ctx context.Context) apibase.Result {
	paramsErr := apibase.NewApiParameterErrors()
	from, err := ctx.URLParamInt("from")
	if err != nil {
		paramsErr.AppendError("$", fmt.Errorf("param from is empty"))
	}
	to, err := ctx.URLParamInt("to")
	if err != nil {
		paramsErr.AppendError("$", fmt.Errorf("param to is empty"))
	}
	targets := ctx.URLParam("targets")
	var targetsList []targetInfo
	if err := json.Unmarshal([]byte(targets), &targetsList); err != nil {
		paramsErr.AppendError("$", fmt.Errorf("param targets format error:%v", err))
	}
	//unit := ctx.URLParam("unit")
	decimal, err := ctx.URLParamInt("decimal")
	if err != nil {
		paramsErr.AppendError("$", fmt.Errorf("param decimal is empty"))
	}
	paramsErr.CheckAndThrowApiParameterErrors()

	var chartInfo ChartsInfo
	// 取第一个target
	if len(targetsList) > 0 {
		var x = make([]float64, 0)
		var y = make([]map[string]interface{}, 0)
		var xEdited = false
		for _, target := range targetsList {
			legendFormat := target.LegendFormat
			err, queryResponse := grafana.GrafanaQuery(target.Expr, from/1000, to/1000, (to/1000-from/1000)/600)
			if err != nil {
				log.Errorf("grafana query error: %v", err)
				return nil
			}
			var maxValue = 0

			resultList := queryResponse.Data.Result
			sort.Sort(resultList)
			if len(resultList) > 0 {
				for index, result := range resultList {
					metric := result.Metric
					if len(metric) == 0 {
						continue
					}
					values := result.Values
					item := map[string]interface{}{}
					item["title"] = formatLegend(legendFormat, metric)
					var data []interface{}
					for _, value := range values {
						if index == 0 && !xEdited {
							x = append(x, value[0].(float64))
						}
						value, _ := FormatFloatCeil(value[1].(string), decimal)
						data = append(data, value)
					}
					if maxValue <= len(values) {
						maxValue = len(values)
					} else {
						zeroSlice := make([]interface{}, 0)
						for i := 0; i < maxValue-len(values); i++ {
							zeroSlice = append(zeroSlice, float64(0))
						}
						data = append(zeroSlice, data...)
					}
					item["data"] = data
					y = append(y, item)

				}
			}
			xEdited = true
		}
		chartInfo.X = x
		chartInfo.Y = y
	}

	return chartInfo
}

func FormatGraphData(from, to, decimal int, targets string) ChartsInfo {
	var targetsList []targetInfo
	if err := json.Unmarshal([]byte(targets), &targetsList); err != nil {
		log.Errorf("param targets format error:%v", err)
	}
	return FormatCharInfo(targetsList, from, to, decimal)
}

func FormatCharInfo(targetsList []targetInfo, from, to, decimal int) ChartsInfo {
	var chartInfo ChartsInfo
	// 取第一个target
	if len(targetsList) > 0 {
		var x = make([]float64, 0)
		var y = make([]map[string]interface{}, 0)
		var xEdited = false
		for _, target := range targetsList {
			legendFormat := target.LegendFormat
			err, queryResponse := grafana.GrafanaQuery(target.Expr, from/1000, to/1000, (to/1000-from/1000)/600)
			if err != nil {
				log.Errorf("grafana query error: %v", err)
			}
			var maxValue = 0

			resultList := queryResponse.Data.Result
			sort.Sort(resultList)
			if len(resultList) > 0 {
				for index, result := range resultList {
					metric := result.Metric
					if len(metric) == 0 {
						continue
					}
					values := result.Values
					item := map[string]interface{}{}
					item["title"] = formatLegend(legendFormat, metric)
					var data []interface{}
					for _, value := range values {
						if index == 0 && !xEdited {
							x = append(x, value[0].(float64))
						}
						value, _ := FormatFloatCeil(value[1].(string), decimal)
						data = append(data, value)
					}
					if maxValue <= len(values) {
						maxValue = len(values)
					} else {
						zeroSlice := make([]interface{}, 0)
						for i := 0; i < maxValue-len(values); i++ {
							zeroSlice = append(zeroSlice, float64(0))
						}
						data = append(zeroSlice, data...)
					}
					item["data"] = data
					y = append(y, item)

				}
			}
			xEdited = true
		}
		chartInfo.X = x
		chartInfo.Y = y
	}

	return chartInfo
}

func formatLegend(legendFormat string, metric map[string]interface{}) string {
	reg := regexp.MustCompile(`\{{(?P<field>\w+)}}`)
	for {
		matches := reg.FindStringSubmatch(legendFormat)
		if matches != nil {
			if value, ok := metric[matches[1]]; ok {
				legendFormat = strings.Replace(legendFormat, matches[0], value.(string), 1)
			}
		} else {
			break
		}
	}
	return legendFormat
}

func FormatFloatCeil(num string, decimal int) (float64, error) {
	value, _ := strconv.ParseFloat(num, 64)
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	res := strconv.FormatFloat(math.Ceil(value*d)/d, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)
}

type GenerateParam struct {
	From int `json:"from"`
	To   int `json:"to"`
}

func StartGenerateReport(ctx context.Context) apibase.Result {
	clusterId, err := GetCurrentClusterId(ctx)
	if err != nil {
		return err
	}
	var param GenerateParam
	if err := ctx.ReadJSON(&param); err != nil {
		log.Errorf("[StartGenerateReport] invalid params")
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("found panic: %v", err)
		}
	}()
	fromDate := time.Unix(int64(param.From/1000), 0).Format("2006-01-02")
	toDate := time.Unix(int64(param.To)/1000, 0).Format("2006-01-02")
	reportName := fmt.Sprintf("运维周报（%s至%s）", fromDate, toDate)
	reportId, err := model.InspectReport.NewInspectReport(reportName, "RUNNING", clusterId)
	if err != nil {
		log.Errorf("[StartGenerateReport] new db record error: %v", err)
		return err
	}
	go generateReport(ctx, param, int(reportId), clusterId)
	return map[string]interface{}{
		"report_id": reportId,
	}
}

func GetReportProgress(ctx context.Context) apibase.Result {
	reportId, err := ctx.URLParamInt("id")
	if err != nil {
		log.Errorf("Get report by id param error: %v", err)
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("found panic: %v", err)
		}
	}()
	reportInfo, err := model.InspectReport.GetById(reportId)
	if err != nil {
		log.Errorf("Get report by id db error: %v", err)
		return err
	}
	return reportInfo
}

func Download(ctx context.Context) apibase.Result {
	paramErrs := apibase.NewApiParameterErrors()
	filePath := ctx.URLParam("file_path")
	if filePath == "" {
		paramErrs.AppendError("$", "缺少文件路径")
	}
	id, err := ctx.URLParamInt("id")
	if err != nil {
		paramErrs.AppendError("$", "缺少id")
	}
	paramErrs.CheckAndThrowApiParameterErrors()
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("open report file error")
		return err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	reportInfo, err := model.InspectReport.GetById(id)
	if err != nil {
		log.Errorf("get report info by id error: %v", err)
		return err
	}
	absolutePath := reportInfo.Name
	fileNames := absolutePath[strings.LastIndex(absolutePath, "/")+1:]
	fileName := url.QueryEscape(fileNames)
	ctx.Header("Content-Disposition", "attachment;filename=\""+fileName+".pdf\"")
	ctx.Write(content)
	return nil
}

func generateReport(ctx context.Context, param GenerateParam, id, clusterId int) error {
	defer func() {
		if err := os.RemoveAll(ImageDir(id)); err != nil {
			log.Errorf("Remove image dir of %d error: %v", id, err)
		}
	}()
	p := util.NewPdfGenerator(id)
	p.AddFont("Simhei", FontFile("simhei.ttf"))
	p.SetFont("Arial", "", 14)
	p.AddFooter()
	p.AddPage()
	// title
	p.AddText(p.PageWidth()*0.4, 18, 7, "Simhei", "巡检报告", 0, 0, 0)
	// 集群状态汇总
	p.AddLine(p.Left())
	p.AddText(p.Left(), 9, 5, "Simhei", "集群状态汇总", 0, 0, 0)
	p.AddText(p.Left(), 6, 2, "Simhei", "1.报告中设计的状态为报告下载时间点的状态", 112, 128, 144)
	p.AddText(p.Left(), 6, 2, "Simhei", "2.状态为“正常”表示节点或应用当前的健康状态为健康，监控指标没有告警。状态为“异常”表示节点或应用当前的健康状态为不健康，或者监控指标有告警", 112, 128, 144)
	p.Ln(3)

	hostStatus := func() {
		hostInfoListInterface := GetHostStatus(ctx)
		if hostInfoList, ok := hostInfoListInterface.([]ReportHostInfo); ok {
			p.AddText(p.Left(), 8, 7, "Simhei", "·节点状态", 0, 0, 0)
			headers := []string{"节点", "cpu", "内存", "系统盘", "数据盘"}
			datas := [][]string{}
			for _, hostInfo := range hostInfoList {
				data := []string{hostInfo.Ip, hostInfo.Cpu, hostInfo.Mem, hostInfo.SystemDisk, hostInfo.DataDisk}
				datas = append(datas, data)
			}
			p.AddTable(datas, headers)
		}
	}

	appStatus := func() {
		p.AddText(p.Left(), 8, 7, "Simhei", "·应用状态", 0, 0, 0)
		serviceStatusMapInterface := GetServiceStatus(ctx)
		if serviceStatusMap, ok := serviceStatusMapInterface.(map[string][]resultLists); ok {
			for k, v := range serviceStatusMap {
				p.AddText(p.Left(), 7, 6, "Simhei", k, 0, 0, 0)
				var headers []string
				datas := [][]string{}
				if k == "DTBase" {
					headers = []string{"服务", "节点", "角色", "状态"}
					for _, info := range v {
						data := []string{info.ServiceName, info.Ip, info.HARole, info.Status}
						datas = append(datas, data)
					}
				} else {
					headers = []string{"服务", "节点", "状态"}
					for _, info := range v {
						data := []string{info.ServiceName, info.Ip, info.Status}
						datas = append(datas, data)
					}
				}
				p.AddTable(datas, headers)
			}
		}
	}

	alertHistory := func() {
		p.AddText(p.Left(), 8, 7, "Simhei", "·告警记录", 0, 0, 0)
		alertHistoryList := FormatAlertList(param.From, param.To)
		headers := []string{"告警名称", "状态", "仪表盘名称(组件)", "仪表盘标题", "告警时间"}
		datas := [][]string{}
		for _, info := range alertHistoryList {
			data := []string{info.AlertName, info.State, info.DashboardName, info.DashboardTitle, info.Time}
			datas = append(datas, data)
		}
		p.AddTable(datas, headers)
	}

	hostStatus()
	appStatus()
	alertHistory()

	err := model.InspectReport.UpdateProgress(id, 30, "", "")

	var configMaps = map[string][]model.BaseTemplateConfig{}
	configList, err := model.InspectReportTemplate.GetTemplateConfig(clusterId)
	for _, config := range configList {
		module := strings.TrimSpace(config.Module)
		moduleConfigList, ok := configMaps[module]
		if !ok {
			moduleConfigList = []model.BaseTemplateConfig{}
		}
		moduleConfigList = append(moduleConfigList, config)
		configMaps[module] = moduleConfigList
	}

	// 集群状态详细内容
	p.AddLine(p.Left())
	p.AddText(p.Left(), 9, 5, "Simhei", "集群状态详细内容", 0, 0, 0)
	p.AddText(p.Left(), 8, 7, "Simhei", "·节点状态", 0, 0, 0)

	if err := os.MkdirAll(ImageDir(p.Id), 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(PdfDir(p.Id), 0755); err != nil {
		return err
	}

	// 渲染节点监控折线图
	hostConfigs, ok := configMaps["System"]
	if !ok {
		log.Errorf("lack host config")
	} else {
		for _, hostConfig := range hostConfigs {
			addLineChart(param.From, param.To, hostConfig, p)
		}
	}

	err = model.InspectReport.UpdateProgress(id, 60, "", "")

	p.AddText(p.Left(), 8, 7, "Simhei", "·应用状态", 0, 0, 0)
	for k, v := range configMaps {
		if k != "System" {
			if existModuleData(param.From, param.To, v) {
				p.AddText(p.Left(), 8, 7, "Simhei", k, 0, 0, 0)
				for _, hostConfig := range v {
					addLineChart(param.From, param.To, hostConfig, p)
				}
			}
		}
	}

	fileStr := Filename(id, strconv.Itoa(id))
	err = p.OutputFileAndClose(fileStr)

	if err != nil {
		log.Errorf("generate pdf error: %v", err)
		err = model.InspectReport.UpdateProgress(id, 60, "", "FAIL")
	}

	err = model.InspectReport.UpdateProgress(id, 100, fileStr, "SUCCESS")
	return err
}

func existModuleData(from, to int, configList []model.BaseTemplateConfig) bool {
	for _, config := range configList {
		chartInfo := FormatGraphData(from, to, config.Decimal, config.Targets)
		if len(chartInfo.X) != 0 {
			return true
		}
	}
	return false
}

func addLineChart(from, to int, hostConfig model.BaseTemplateConfig, p *util.PdfGenerator) {
	chartsInfo := FormatGraphData(from, to, hostConfig.Decimal, hostConfig.Targets)
	series := make([]chart.Series, len(chartsInfo.Y))
	if len(chartsInfo.X) == 0 {
		return
	}
	for i := 0; i < len(chartsInfo.Y); i++ {
		xValues := make([]time.Time, len(chartsInfo.X))
		yValues := make([]float64, len(chartsInfo.X))
		title := chartsInfo.Y[i]["title"].(string)
		datas := chartsInfo.Y[i]["data"].([]interface{})
		for j := 0; j < len(chartsInfo.X); j++ {
			xValues[j] = time.Unix(int64(chartsInfo.X[j]), 0)
			if hostConfig.Unit == "byte" {
				yValues[j] = datas[j].(float64) / (1000 * 1000 * 1000)
			} else {
				yValues[j] = datas[j].(float64)
			}
		}
		series[i] = chart.TimeSeries{
			Name: title,
			Style: chart.Style{
				StrokeColor: drawing.Color{
					R: uint8(rand.Intn(256)),
					G: uint8(rand.Intn(256)),
					B: uint8(rand.Intn(256)),
					A: uint8(256 - 1),
				},
			},
			XValues: xValues,
			YValues: yValues,
		}
	}
	lineChartStyle := chart.Style{
		Padding: chart.Box{
			Top: 40,
		},
	}
	ya := defineYaxis(hostConfig.Metric)
	graph := chart.Chart{
		Title: hostConfig.Metric,
		TitleStyle: chart.Style{
			Font: GetChineseFont(),
		},
		Background: lineChartStyle,
		XAxis: chart.XAxis{
			Name: "时间",
			NameStyle: chart.Style{
				Font: GetChineseFont(),
			},
			ValueFormatter: chart.TimeValueFormatterWithFormat("2006-01-02 15:04:05"),
		},
		YAxis:  ya,
		Series: series,
	}
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}
	imgName := ImageFile(p.Id, hostConfig.Module+"_"+strings.ReplaceAll(hostConfig.Metric, "/", "_")+".png")
	f, _ := os.Create(imgName)
	defer f.Close()
	fmt.Println(imgName)
	graph.Render(chart.PNG, f)
	p.AddLineChart("png", imgName)
}

func defineYaxis(metricName string) chart.YAxis {
	ya := chart.YAxis{
		Name: "值",
		NameStyle: chart.Style{
			Font: GetChineseFont(),
		},
	}
	var max float64
	if strings.Contains(metricName, "up") || strings.Contains(metricName, "Up") {
		max = 1
	} else if strings.Contains(metricName, "%") {
		max = 100
	} else {
		max = 10
	}
	ya.Range = &chart.ContinuousRange{
		Max: max,
	}
	return ya
}

func GetChineseFont() *truetype.Font {
	fontBytes, err := ioutil.ReadFile(FontFile("simhei.ttf"))
	if err != nil {
		return nil
	}
	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil
	}
	return font
}
