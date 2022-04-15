package enums

import "reflect"

/*
 @Author: zhijian
 @Date: 2021/5/28 10:28
 @Description: 操作对象枚举
*/

type operationObjType struct {
	Product EnumValueType
	Svc     EnumValueType
	Host    EnumValueType
}

func (c operationObjType) List() (enumValues []EnumValueType) {
	v := reflect.ValueOf(c)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		enumValues = append(enumValues, field.Interface().(EnumValueType))
	}
	return enumValues
}

var OperationObjType = operationObjType{
	Product: EnumValueType{
		Code: 1,
		Desc: "产品包",
	},
	Svc: EnumValueType{
		Code: 2,
		Desc: "服务",
	},
	Host: EnumValueType{
		Code: 3,
		Desc: "主机",
	},
}
