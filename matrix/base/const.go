package base

import (
	"time"
)

const (
	TsLayout = "2006-01-02 15:04:05"
	WebRoot  = "./easyagent"

	StatusInit    int = -1
	StatusRunning int = 0
	StatusStopped int = 1
	StatusError   int = 2

	SwitchOn  string = "on"
	SwitchOff string = "off"
)

var AgentInstallStateName = map[int]string{
	1:  "管控安装成功",
	-1: "管控安装失败",
	2:  "script安装成功",
	-2: "script安装失败",
	3:  "主机初始化成功",
	-3: "主机初始化失败",
}
var SupportClusterType = map[string]struct{}{
	"kubernetes": {},
	"hosts":      {},
}

type Time time.Time

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in TsLayout format, with sub-second precision added if present.
func (this Time) MarshalJSON() ([]byte, error) {
	t := time.Time(this)

	b := make([]byte, 0, len(TsLayout)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, TsLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in TsLayout format.
func (this *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	t, err := time.Parse(`"`+TsLayout+`"`, string(data))
	this = (*Time)(&t)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in TsLayout format, with sub-second precision added if present.
func (this Time) MarshalText() ([]byte, error) {
	t := time.Time(this)

	b := make([]byte, 0, len(TsLayout))
	return t.AppendFormat(b, TsLayout), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in TsLayout format.
func (this *Time) UnmarshalText(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	t, err := time.Parse(TsLayout, string(data))
	this = (*Time)(&t)
	return err
}
