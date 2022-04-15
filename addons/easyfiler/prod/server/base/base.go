package base

var (
	VERSION      = "dev-snapshot"
	_SYSTEM_FAIL = make(chan SystemFailure)
)

func ConfigureProductVersion(v string) {
	VERSION = v
}
