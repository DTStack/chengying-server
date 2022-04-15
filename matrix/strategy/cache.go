package strategy

import (
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"sync"
)

type safeStrategyCacheMap struct {
	sync.RWMutex
	s map[int]*model.StrategyInfo
	r map[int]*model.StrategyResourceInfo
	a map[int][]*model.StrategyAssignInfo
}

var (
	StrategyCacheMap = &safeStrategyCacheMap{
		s: make(map[int]*model.StrategyInfo),
		r: make(map[int]*model.StrategyResourceInfo),
		a: make(map[int][]*model.StrategyAssignInfo),
	}
)

func (this *safeStrategyCacheMap) ReInit(
	s map[int]*model.StrategyInfo,
	r map[int]*model.StrategyResourceInfo,
	a map[int][]*model.StrategyAssignInfo) {

	this.Lock()
	defer this.Unlock()
	this.s = s
	this.r = r
	this.a = a
}

func (this *safeStrategyCacheMap) GetStrategy() map[int]*model.StrategyInfo {
	this.RLock()
	defer this.RUnlock()
	return this.s
}

func (this *safeStrategyCacheMap) GetResource() map[int]*model.StrategyResourceInfo {
	this.RLock()
	defer this.RUnlock()
	return this.r
}

func (this *safeStrategyCacheMap) GetAssign() map[int][]*model.StrategyAssignInfo {
	this.RLock()
	defer this.RUnlock()
	return this.a
}
