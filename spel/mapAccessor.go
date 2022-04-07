package spel

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : mapAccessor.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:16
* 修改历史 : 1. [2022/4/7 18:16] 创建文件 by NST
*/
import "reflect"

type MapAccessor struct {
}

func (this MapAccessor) CanRead(context EvaluationContext, target interface{}, name string) bool {
	m, ok := target.(map[string]interface{})
	return ok && m[name] != nil
}

func (this MapAccessor) Read(context EvaluationContext, target interface{}, name string) TypedValue {
	m, ok := target.(map[string]interface{})
	if !ok {
		panic("Target must be of type Map")
	}
	value := m[name]
	if value == nil {
		panic("Map does not contain a value for key")
	}
	return TypedValue{Value: value}
}

func (this MapAccessor) GetSpecificTargetClasses() interface{} {
	return reflect.Map
}
