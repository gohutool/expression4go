package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : literalValue.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:34
* 修改历史 : 1. [2022/4/7 19:34] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

type LiteralValue struct {
	*SpelNodeImpl
}

func (l *LiteralValue) GetLiteralValue() TypedValue {
	return TypedValue{}
}
