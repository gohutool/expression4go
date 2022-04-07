package support

import "github.com/gohutool/expression4go/spel"

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : booleanTypedValue.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:28
* 修改历史 : 1. [2022/4/7 18:28] 创建文件 by NST
*/

type BooleanTypedValue struct {
	*spel.TypedValue
}

func (b *BooleanTypedValue) ForValue(bool2 bool) spel.TypedValue {
	boo := spel.TypedValue{}
	if bool2 {
		boo.Value = true
		return boo
	}

	boo.Value = false
	return boo
}
