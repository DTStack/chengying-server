package options

import (
	"github.com/spf13/pflag"
	"sync"
)

var opt *Options
var once sync.Once

const (
	DefaultLogDebug         = false
	DefaultLogTiemLayout    = "2006-01-02 15:04:05"
	DefaultElectionLockName = "dtstack-operator"
)

//option will be used
type Options struct {
	LogDebug         bool
	LogTime          string
	WatchNamespace   string
	ElectionLockName string
}

func GetOptions() *Options {
	once.Do(func() {
		opt = &Options{
			LogDebug:         DefaultLogDebug,
			LogTime:          DefaultLogTiemLayout,
			ElectionLockName: DefaultElectionLockName,
		}
	})
	return opt
}

func (opt *Options) AddToFlagSet(fs *pflag.FlagSet) {
	fs.BoolVar(&opt.LogDebug, "log-debug", DefaultLogDebug, "set logger in debug mode")
	fs.StringVar(&opt.LogTime, "log-time", DefaultLogTiemLayout, "set log time layout format")
	fs.StringVar(&opt.WatchNamespace, "watch-namespace", "", "namespace the operator is watched")
	fs.StringVar(&opt.ElectionLockName, "election-lock-name", DefaultElectionLockName, "name of configmap lock name")
}
