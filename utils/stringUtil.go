package utils

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : stringUtil.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:08
* 修改历史 : 1. [2022/4/7 18:08] 创建文件 by NST
*/

func IndexOf(expressionString string, ch string, fromIndex int) int {
	max := len(expressionString)
	if fromIndex < 0 {
		fromIndex = 0
	}
	if fromIndex >= max {
		return -1
	}
	for i := fromIndex; i < max; i++ {
		if string(expressionString[i]) == ch {
			return i
		}
	}
	return -1
}

func BinarySearch(a []string, v interface{}) int {
	for i, i2 := range a {
		if i2 == v {
			return i
		}
	}
	return -1
}
