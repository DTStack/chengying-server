package strategy

import (
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"time"
)

var (
	Duration = 60
)

func SyncStrategies() {
	duration := time.Duration(Duration) * time.Second
	for {
		log.Infof("SyncStrategies ...")
		syncStrategies()
		time.Sleep(duration)
	}
}

func syncStrategies() {
	err, strategy := model.StrategyList.GetDeployedStrategyList()
	if err != nil {
		log.Errorf("err :%v", err.Error())
		return
	}
	err, strategyResources := model.StrategyResourceList.GetStrategyResourceList()
	if err != nil {
		log.Errorf("err :%v", err.Error())
		return
	}
	err, strategyAssigns := model.StrategyAssignList.GetStrategyAssignList()
	if err != nil {
		log.Errorf("err :%v", err.Error())
		return
	}
	rebuildStrategyMap(strategy, strategyResources, strategyAssigns)
}

func rebuildStrategyMap(strategys []*model.StrategyInfo,
	resources []*model.StrategyResourceInfo,
	assing []*model.StrategyAssignInfo) {
	s := make(map[int]*model.StrategyInfo)
	r := make(map[int]*model.StrategyResourceInfo)
	a := make(map[int][]*model.StrategyAssignInfo)

	for _, ss := range strategys {
		s[ss.ID] = ss
	}
	for _, rr := range resources {
		r[rr.StrategyId] = rr
	}
	for _, aa := range assing {
		a[aa.StrategyId] = append(a[aa.StrategyId], aa)
	}
	StrategyCacheMap.ReInit(s, r, a)
}
