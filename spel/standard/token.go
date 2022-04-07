package standard

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : token.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:29
* 修改历史 : 1. [2022/4/7 18:29] 创建文件 by NST
*/

type Token struct {
	Kind     TokenKind
	Data     string
	StartPos int
	EndPos   int
}

func (t Token) isNumericRelationalOperator() bool {
	return true
}

func (t Token) StringValue() string {
	if t.Data != "" {
		return t.Data
	}
	return ""
}

func (t Token) IsNumericRelationalOperator() bool {
	return t.Kind.TokenKindType == GT || t.Kind.TokenKindType == GE || t.Kind.TokenKindType == LT ||
		t.Kind.TokenKindType == LE || t.Kind.TokenKindType == EQ || t.Kind.TokenKindType == NE
}
