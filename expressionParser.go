package expression4go

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : expressionParser.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:24
* 修改历史 : 1. [2022/4/7 18:24] 创建文件 by NST
*/

type ExpressionParser interface {
	ParseExpression(var1 string) Expression

	DoParseExpression(var1 string) Expression
}
