package standard

import . "github.com/gohutool/expression4go/spel"
import . "github.com/gohutool/expression4go/spel/ast"

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : spelExpression.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:31
* 修改历史 : 1. [2022/4/7 18:31] 创建文件 by NST
*/

type SpelExpression struct {
	Expression        string
	Configuration     SpelParserConfiguration
	EvaluationContext EvaluationContext
	Ast               SpelNode
}

func (e SpelExpression) GetExpressionString() string {
	return e.Expression
}

func (e *SpelExpression) GetValue() interface{} {
	context := e.getEvaluationContext()
	state := ExpressionState{RelatedContext: context}
	return e.Ast.GetValue(state)
}

func (e *SpelExpression) GetValueContext(context EvaluationContext) interface{} {
	if context == nil {
		panic("EvaluationContext is required")
	}
	state := ExpressionState{RelatedContext: context}
	typedResultValue := e.Ast.GetValueInternal(state)
	return typedResultValue.Value
}

func (e *SpelExpression) getEvaluationContext() EvaluationContext {
	if e.EvaluationContext == nil {
		context := StandardEvaluationContext{}
		e.EvaluationContext = &context
	}
	return e.EvaluationContext
}
