package examples

import (
	"fmt"
	"github.com/gohutool/expression4go"
	"github.com/gohutool/expression4go/spel"
	"testing"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : example_test.go
* 文件路径 : examples
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 20:04
* 修改历史 : 1. [2022/4/7 20:04] 创建文件 by NST
*/

func TestExample(t *testing.T) {
	context := spel.StandardEvaluationContext{}
	m := make(map[string]interface{})
	m["name"] = "lisi"
	m["age"] = 18
	context.SetVariables(m)
	parser := expression4go.SpelExpressionParser{}
	expressionString := "#name=='lisi'"
	//expressionString := "#name" //返回lisi
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)

	fmt.Println(valueContext)
}
