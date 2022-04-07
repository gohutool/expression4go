package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : opOr.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:50
* 修改历史 : 1. [2022/4/7 19:50] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
	. "github.com/gohutool/expression4go/support"
)

type OpOr struct {
	*Operator
}

func (o *OpOr) GetValueInternal(expressionState ExpressionState) TypedValue {
	if getBooleanValue(expressionState, o.getLeftOperand()) {
		value := BooleanTypedValue{}
		return value.ForValue(true)
	}
	booleanValue := getBooleanValue(expressionState, o.getRightOperand())
	value := BooleanTypedValue{}
	return value.ForValue(booleanValue)
}
