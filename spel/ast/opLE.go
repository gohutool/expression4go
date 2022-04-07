package ast

import "github.com/gohutool/expression4go/support"

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : opLE.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:39
* 修改历史 : 1. [2022/4/7 19:39] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

////处理小于等于
type OpLE struct {
	*Operator
}

func (o *OpLE) GetValueInternal(expressionState ExpressionState) TypedValue {
	value := support.BooleanTypedValue{}
	left := o.getLeftOperand().GetValueInternal(expressionState).Value
	right := o.getRightOperand().GetValueInternal(expressionState).Value
	checkType := checkType(left, right)
	if !checkType {
		return value.ForValue(checkType)
	}
	o.leftActualDescriptor = o.toDescriptorFromObject(left)
	o.rightActualDescriptor = o.toDescriptorFromObject(right)
	var check bool
	leftV, ok := left.(int)
	if ok {
		rightV := right.(int)
		check = leftV <= rightV
	} else {
		leftV, ok := left.(float64)
		if ok {
			rightV := right.(float64)
			check = leftV <= rightV
		}
	}
	return value.ForValue(check)
}
