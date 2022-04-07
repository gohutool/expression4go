package expression4go

import "github.com/gohutool/expression4go/spel"

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : expression.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:10
* 修改历史 : 1. [2022/4/7 18:10] 创建文件 by NST
*/

//获取值
type Expression interface {
	GetExpressionString() string

	GetValue() interface{}

	GetValueContext(context spel.EvaluationContext) interface{}
}
