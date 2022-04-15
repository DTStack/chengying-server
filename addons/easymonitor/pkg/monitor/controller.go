package monitor

type Controller interface {
	Add()
	Run(headiness int, stopCh chan struct{})
}
