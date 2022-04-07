package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : variableReference.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:54
* 修改历史 : 1. [2022/4/7 19:54] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

//map参数
const (
	THIS = "this"
	ROOT = "root"
)

type VariableReference struct {
	*SpelNodeImpl
	Name string
}

func (v VariableReference) GetValueInternal(state ExpressionState) TypedValue {
	return state.LookupVariable(v.Name)
}
