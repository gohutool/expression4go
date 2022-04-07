package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : compoundExpression.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:38
* 修改历史 : 1. [2022/4/7 18:38] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

type CompoundExpression struct {
	*SpelNodeImpl
}

//复合表达式
func (this *CompoundExpression) GetValueInternal(state ExpressionState) TypedValue {
	ref := this.GetValueRef(state)
	return ref.GetValue()

}
func (this CompoundExpression) GetValueRef(state ExpressionState) ValueRef {
	if this.getChildCount() == 1 {
		return this.Children[0].GetValueRef(state)
	}
	nextNode := this.Children[0]
	result := nextNode.GetValueInternal(state)
	count := this.getChildCount()
	for i := 1; i < count-1; i++ {
		defer state.PopActiveContextObject()
		state.PushActiveContextObject(result)
		nextNode = this.Children[i]
		result = nextNode.GetValueInternal(state)
	}
	defer state.PopActiveContextObject()
	state.PushActiveContextObject(result)
	nextNode = this.Children[count-1]
	return nextNode.GetValueRef(state)
}

func (o *CompoundExpression) getChildCount() int {
	return len(o.Children)
}
