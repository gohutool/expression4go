package spel

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : reflectivePropertyAccessor.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:18
* 修改历史 : 1. [2022/4/7 18:18] 创建文件 by NST
*/

import (
	"reflect"
	"unsafe"
)

//反射处理字段
type ReflectivePropertyAccessor struct {
}

func (this ReflectivePropertyAccessor) CanRead(context EvaluationContext, target interface{}, name string) bool {
	if target == nil {
		return false
	}
	kind := reflect.TypeOf(target).Kind()
	if (kind == reflect.Slice || kind == reflect.Array) && name == "len" {
		return true
	}
	return false
}

func (this ReflectivePropertyAccessor) Read(context EvaluationContext, target interface{}, name string) TypedValue {
	r, ok := target.(reflect.Value)
	var value interface{}
	if ok {
		typeOfCat := r.Type()
		byName, _ := typeOfCat.FieldByName(name)
		valueType := byName.Type.Kind()
		v := r.FieldByName(name)
		switch valueType {
		case reflect.Int:
			//转为Int64
			valueInt := v.Int()
			// 将 int64 转化为 int
			value = *(*int)(unsafe.Pointer(&valueInt))
			break
		case reflect.String:
			value = v.String()
			break
		case reflect.Float64:
			value = v.Float()
			break
		case reflect.Struct:
			value = v
			break
		}
	}
	return TypedValue{Value: value}
}

func (this ReflectivePropertyAccessor) GetSpecificTargetClasses() interface{} {
	return nil
}
