package events

const (
	SEP = "_"
)

type Eventer interface {
	Info() string
}
