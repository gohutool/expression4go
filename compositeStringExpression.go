package expression4go

import "github.com/gohutool/expression4go/spel"

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : compositeStringExpression.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:23
* 修改历史 : 1. [2022/4/7 18:23] 创建文件 by NST
*/

type CompositeStringExpression struct {
	*spel.ExpressionImpl
	ExpressionString string
	Expressions      []Expression
}

func (c *CompositeStringExpression) GetExpressionString() string {
	return c.ExpressionString
}

func (c *CompositeStringExpression) GetValue() interface{} {
	//s := ""

	return "c.literalValue"
}
