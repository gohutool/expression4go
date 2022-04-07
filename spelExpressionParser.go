package expression4go

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : spelExpressionParser.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:05
* 修改历史 : 1. [2022/4/7 18:05] 创建文件 by NST
*/

type SpelExpressionParser struct {
	*TemplateAwareExpressionParser
}

func (s *TemplateAwareExpressionParser) DoParseExpression(expressionString string) Expression {
	parser := InternalSpelExpressionParser{}
	return parser.DoParseExpression(expressionString)
}
