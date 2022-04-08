package examples

import (
	"fmt"
	. "github.com/gohutool/expression4go"
	. "github.com/gohutool/expression4go/spel"
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

var (
	context = StandardEvaluationContext{}
	m       = make(map[string]interface{})
	parser  = SpelExpressionParser{}
)

type Order struct {
	name string
	num  int
}

func TestExample(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m["name"] = "expression4go"
	m["age"] = 1

	//
	//map1 := make(map[string]interface{})
	//m["map"] = map1
	//map1["name"] = "davidliu"
	//map1["age"] = 10

	context.SetVariables(m)
	parser := SpelExpressionParser{}
	expressionString := "#{name}"
	//expressionString := "#name=='lisi'"
	//expressionString := "#name" //返回lisi
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)

	fmt.Println(valueContext)
}

func TestStruct(t *testing.T) {
	//context.AddPropertyAccessor(MapAccessor{})
	m1 := make(map[string]interface{})

	order := Order{
		name: "expression4go",
		num:  1000}
	m1["order"] = order
	m["data"] = order

	context.SetVariables(m)
	expressionString := "#data"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为", valueContext)
}

func TestCompound2(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m["name"] = "expression4go"
	context.SetVariables(m)
	expressionString := "${name}"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为", valueContext)
}

func TestCompound(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m["name"] = "expression4go"
	m["age"] = 1
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m2["num"] = 12
	m1["code"] = m2
	m["order"] = m1
	context.SetVariables(m)
	expressionString := "${order.code.num}"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为", valueContext)
}

func TestIndex(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m1 := make(map[string]interface{})
	m["name"] = "expression4go"
	m["age"] = 1
	//切片
	//orders := make([]Order, 2)
	//数组
	orders := [2]Order{}
	orders[0] = Order{name: "expression4go-1", num: 12}
	orders[1] = Order{name: "expression4go-2", num: 24}
	m1["code"] = orders
	m["order"] = m1
	context.SetVariables(m)
	expressionString := "${order.code[0].name}"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为", valueContext)
}

func TestSlice(t *testing.T) {
	var arr []any
	arr = append(arr, "liuyong")
	fmt.Println(arr...)
}

func TestIndex2(t *testing.T) {
	context.AddPropertyAccessor(MapAccessor{})
	m1 := make(map[string]interface{})
	m["name"] = "expression4go"
	m["age"] = 1
	//切片
	//orders := make([]Order, 2)
	//数组
	orders := [2]int{}
	orders[0] = 1
	orders[1] = 2
	m1["code"] = orders
	m["order"] = m1
	context.SetVariables(m)
	expressionString := "${order.code[0]}"
	valueContext := parser.ParseExpression(expressionString).GetValueContext(&context)
	fmt.Println("结果为", valueContext)
}
