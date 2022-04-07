package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : spelNode.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:52
* 修改历史 : 1. [2022/4/7 19:52] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

//SpelNode
//表达式对象
type SpelNode interface {
	GetValue(expressionState ExpressionState) interface{}

	GetValueInternal(expressionState ExpressionState) TypedValue

	GetValueRef(state ExpressionState) ValueRef

	GetStartPosition() int

	GetEndPosition() int
}
