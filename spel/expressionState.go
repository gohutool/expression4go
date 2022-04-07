package spel

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : expressionState.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:15
* 修改历史 : 1. [2022/4/7 18:15] 创建文件 by NST
*/
import "container/list"

//根据KEY获取MAP的value
type ExpressionState struct {
	RelatedContext EvaluationContext

	RootObject TypedValue

	ContextObjects list.List

	VariableScopes list.List
}

func (this *ExpressionState) LookupVariable(name string) TypedValue {
	variable := this.RelatedContext.LookupVariable(name)
	return TypedValue{Value: variable}
}

func (this *ExpressionState) PopActiveContextObjectNull() {
	front := this.ContextObjects.Front()
	this.ContextObjects.Remove(front)
}

func (this *ExpressionState) PushActiveContextObject(obj TypedValue) {
	this.ContextObjects.PushFront(obj)
}

func (this *ExpressionState) PopActiveContextObject() {
	front := this.ContextObjects.Front()
	this.ContextObjects.Remove(front)
}

func (this *ExpressionState) GetActiveContextObject() TypedValue {
	return this.ContextObjects.Front().Value.(TypedValue)
}

func (this *ExpressionState) GetEvaluationContext() EvaluationContext {
	return this.RelatedContext
}

func (this *ExpressionState) GetRootContextObject() TypedValue {
	return this.RootObject
}
