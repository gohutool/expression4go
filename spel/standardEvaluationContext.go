package spel

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : standardEvaluationContext.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:19
* 修改历史 : 1. [2022/4/7 18:19] 创建文件 by NST
*/

//map存储
type StandardEvaluationContext struct {
	Variables map[string]interface{}

	propertyAccessors []PropertyAccessor
}

func (this *StandardEvaluationContext) AddPropertyAccessor(resolver PropertyAccessor) {
	this.AddBeforeDefault(this.initPropertyAccessors(), resolver)
}

func (this *StandardEvaluationContext) AddBeforeDefault(resolvers []PropertyAccessor, resolver PropertyAccessor) {
	resolvers = append(resolvers, resolver)
	this.propertyAccessors = resolvers
}
func (this *StandardEvaluationContext) SetVariable(var1 string, var2 map[string]interface{}) {
	this.Variables[var1] = var2
}

func (this *StandardEvaluationContext) SetVariables(var2 map[string]interface{}) {
	this.Variables = var2
}

func (this StandardEvaluationContext) LookupVariable(name string) interface{} {
	return this.Variables[name]
}

func (this StandardEvaluationContext) GetPropertyAccessors() []PropertyAccessor {
	return this.initPropertyAccessors()
}

func (this StandardEvaluationContext) initPropertyAccessors() []PropertyAccessor {
	accessors := this.propertyAccessors
	if accessors == nil {
		accessors = make([]PropertyAccessor, 1)
		accessors[0] = ReflectivePropertyAccessor{}
		this.propertyAccessors = accessors

	}
	return accessors
}
