package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : spelNodeImpl.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:53
* 修改历史 : 1. [2022/4/7 19:53] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
)

type SpelNodeImpl struct {
	Children           []SpelNode
	Parent             SpelNode
	Pos                int
	exitTypeDescriptor string
}

func (this SpelNodeImpl) GetValueInternal(expressionState ExpressionState) TypedValue {
	return this.GetValueInternal(expressionState)
}

func (this SpelNodeImpl) GetValue(expressionState ExpressionState) interface{} {
	return this.GetValueInternal(expressionState)
}

func (this SpelNodeImpl) GetValueRef(state ExpressionState) ValueRef {
	return nil
}
func (this SpelNodeImpl) GetStartPosition() int {
	return this.Pos >> 16
}

func (this SpelNodeImpl) GetEndPosition() int {
	return this.Pos & 0xffff
}
