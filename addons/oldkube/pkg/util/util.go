package util

import (
	"fmt"
)

func MultiSizeConvert(size1, size2 int64) (string, string) {
	sizeUnits := [...]string{"B", "KB", "MB", "GB", "TB"}
	f1 := float32(size1)
	f2 := float32(size2)
	for _, v := range sizeUnits {
		if f1 < 1024 && f2 < 1024 {
			return fmt.Sprintf("%.2f"+v, f1), fmt.Sprintf("%.2f"+v, f2)
		} else {
			f1 = f1 / 1024
			f2 = f2 / 1024
		}
	}
	return fmt.Sprintf("%.2f"+sizeUnits[len(sizeUnits)-1], f1), fmt.Sprintf("%.2f"+sizeUnits[len(sizeUnits)-1], f1)
}

func MapConvert(m map[interface{}]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range m {
		switch valueType := v.(type) {
		case map[interface{}]interface{}:
			result[fmt.Sprint(k)] = MapConvert(valueType)
		default:
			result[fmt.Sprint(k)] = v
		}
	}
	return result
}
