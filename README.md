# expression4go
a expression launage toolkit for golang like as expression4j

![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)


# Introduce
This project is come from an unmaintained fork, the branch can see https://github.com/heartlhj/go-expression

- left only so it doesn't break imports.
- some enhancement if to need

# Feature
This project's idea is come from expression4j, and support a expression lanugage toolkit for golang.

- Support system output like stdout, stderr and so on


# Usage
- Add expression4go with the following import

```
import "github.com/gohutool/expression4go"
```

- Support standard expression

```
var (
	context = StandardEvaluationContext{}
	m       = make(map[string]interface{})
	parser  = SpelExpressionParser{}
)

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
```

- Support Map Object expression

```
var (
	context = StandardEvaluationContext{}
	m       = make(map[string]interface{})
	parser  = SpelExpressionParser{}
)

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
```

- Support Struct Object expression
```
var (
	context = StandardEvaluationContext{}
	m       = make(map[string]interface{})
	parser  = SpelExpressionParser{}
)

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
```


- Support Slice Object expression
```
var (
	context = StandardEvaluationContext{}
	m       = make(map[string]interface{})
	parser  = SpelExpressionParser{}
)

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
```