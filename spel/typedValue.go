package spel

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : typedValue.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:12
* 修改历史 : 1. [2022/4/7 18:12] 创建文件 by NST
*/

import "reflect"

type TypedValue struct {
	Value interface{}

	typeDescriptor TypeDescriptor
}

func (this TypedValue) GetTypeDescriptor() TypeDescriptor {
	if reflect.DeepEqual(this.typeDescriptor, TypeDescriptor{}) && this.Value == nil {
		object, _ := TypeDescriptor{}.ForObject(this.Value)
		this.typeDescriptor = object
	}
	return this.typeDescriptor
}
