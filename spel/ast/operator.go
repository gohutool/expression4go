package ast

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : operator.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 19:50
* 修改历史 : 1. [2022/4/7 19:50] 创建文件 by NST
*/

import "reflect"

type Operator struct {
	*SpelNodeImpl
	operatorName          string
	leftActualDescriptor  string
	rightActualDescriptor string
}

func (s *SpelNodeImpl) getLeftOperand() SpelNode {
	return s.Children[0]
}

func (s *SpelNodeImpl) getRightOperand() SpelNode {
	return s.Children[1]
}

func (s *SpelNodeImpl) toDescriptorFromObject(value interface{}) string {
	return reflect.TypeOf(value).String()
}

func (s *SpelNodeImpl) equalityCheck(left interface{}, right interface{}) bool {
	if s.toDescriptorFromObject(left) == s.toDescriptorFromObject(right) {
		if left == right {
			return true
		}
	}
	return false
}

func checkType(left interface{}, right interface{}) bool {
	if reflect.TypeOf(left).String() != reflect.TypeOf(right).String() {
		return false
	}
	return true
}
