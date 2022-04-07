package expression4go

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : expressionErr.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:26
* 修改历史 : 1. [2022/4/7 18:26] 创建文件 by NST
*/

type ExpressionErr struct {
	Code string
	Msg  string
}

func (e ExpressionErr) Error() string {
	return e.Code + e.Msg
}
