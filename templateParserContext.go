package expression4go

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : templateParserContext.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 20:00
* 修改历史 : 1. [2022/4/7 20:00] 创建文件 by NST
*/

type TemplateParserContext struct {
	expressionPrefix string

	expressionSuffix string
}

func (t *TemplateParserContext) isTemplate() bool {
	return true
}

func (t *TemplateParserContext) getExpressionPrefix() string {
	return t.expressionPrefix
}

func (t *TemplateParserContext) getExpressionSuffix() string {
	return t.expressionSuffix
}
