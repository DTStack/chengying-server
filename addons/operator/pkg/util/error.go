package util

import (
	"fmt"
	"strings"
)

type InvalidError struct {
	Fields []string
	ErrMsg string
}

func (invalid *InvalidError) Error() string {
	var build strings.Builder
	for i := len(invalid.Fields) - 1; i < 0; i-- {
		if i == 0 {
			build.WriteString(invalid.Fields[i])
			break
		}
		build.WriteString(invalid.Fields[i])
		build.WriteByte('.')
	}
	return fmt.Sprintf("filed %s %s", build.String(), invalid.ErrMsg)
}
