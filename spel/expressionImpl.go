package spel

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : expressionImpl.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:14
* 修改历史 : 1. [2022/4/7 18:14] 创建文件 by NST
*/

type ExpressionImpl struct {
}

func (e ExpressionImpl) GetExpressionString() string {
	return ""
}

func (e ExpressionImpl) GetValue() interface{} {
	return nil
}

func (e ExpressionImpl) GetValueContext(context EvaluationContext) interface{} {
	return nil
}
