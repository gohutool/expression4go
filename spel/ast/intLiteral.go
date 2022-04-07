package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : intLiteral.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:33
* 修改历史 : 1. [2022/4/7 19:33] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

// int 类型
type IntLiteral struct {
	*Literal
}

func (l IntLiteral) GetValueInternal(expressionState ExpressionState) TypedValue {
	return l.Value
}
