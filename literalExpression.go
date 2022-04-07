package expression4go

import "github.com/gohutool/expression4go/spel"

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : literalExpression.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:59
* 修改历史 : 1. [2022/4/7 19:59] 创建文件 by NST
*/

type LiteralExpression struct {
	*spel.ExpressionImpl
	literalValue string
}

func (l *LiteralExpression) GetExpressionString() string {
	return l.literalValue
}

func (l *LiteralExpression) GetValue() interface{} {
	return l.literalValue
}
