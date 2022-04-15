package kube

// db table related to k8s
type KubeTable interface {
	Prepare() error
}

type PrepareFunc func() error

func (p PrepareFunc) Prepare() error{
	return p()
}
