package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : opNE.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:49
* 修改历史 : 1. [2022/4/7 19:49] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

//不等于
type OpNE struct {
	*Operator
}

func (o *OpNE) GetValueInternal(expressionState ExpressionState) TypedValue {
	left := o.getLeftOperand().GetValueInternal(expressionState).Value
	right := o.getRightOperand().GetValueInternal(expressionState).Value
	o.leftActualDescriptor = o.toDescriptorFromObject(left)
	o.rightActualDescriptor = o.toDescriptorFromObject(right)
	check := !o.equalityCheck(left, right)
	value := BooleanTypedValue{}
	return value.ForValue(check)
}
