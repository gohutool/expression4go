package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : opAnd.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:36
* 修改历史 : 1. [2022/4/7 19:36] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

//与操作 eg:  "#name=='lisi' && #age>=18"
type OpAnd struct {
	*Operator
}

func (o *OpAnd) GetValueInternal(expressionState ExpressionState) TypedValue {
	if !getBooleanValue(expressionState, o.getLeftOperand()) {
		value := BooleanTypedValue{}
		return value.ForValue(false)
	}
	booleanValue := getBooleanValue(expressionState, o.getRightOperand())
	value := BooleanTypedValue{}
	return value.ForValue(booleanValue)
}

func getBooleanValue(state ExpressionState, operand SpelNode) bool {
	value := operand.GetValueInternal(state)
	if value.Value == nil {
		panic("Type conversion problem, cannot convert from [null] to bool")
	}
	return value.Value.(bool)
}
