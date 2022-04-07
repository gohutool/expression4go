package standard

import (
	"fmt"
	"github.com/gohutool/expression4go/utils"
	"strings"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : tokenizer.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:29
* 修改历史 : 1. [2022/4/7 18:29] 创建文件 by NST
*/
//将EL表达一些输入数据转换成令牌流，然后可以对其进行解析
const (
	IS_DIGIT    = 0x01
	IS_HEXDIGIT = 0x02
	IS_ALPHA    = 0x04
)

var FLAGS = [256]rune{}

type Tokenizer struct {
	ExpressionString string
	charsToProcess   []rune
	pos              int
	max              int
	tokens           []Token
}

func (t *Tokenizer) InitTokenizer() {
	t.initFlags()
	expressionString := t.initExpression(t.ExpressionString)
	runes := []rune(expressionString)
	t.charsToProcess = runes
	t.max = len(t.charsToProcess)
	t.pos = 0
}

func (t *Tokenizer) initExpression(expressionString string) string {
	ok := strings.Contains(expressionString, "{")
	if ok {
		left := strings.Contains(expressionString, "}")
		if !left {
			panic("Missing closing '}'")
		}
		index := strings.LastIndex(expressionString, "}")
		if index == (len(expressionString) - 1) {
			return expressionString[0:index]
		}
	}
	return expressionString
}

func (t *Tokenizer) initFlags() {

	for ch := '0'; ch <= '9'; ch++ {
		FLAGS[ch] |= IS_DIGIT | IS_HEXDIGIT
	}
	for ch := 'A'; ch <= 'F'; ch++ {
		FLAGS[ch] |= IS_DIGIT | IS_HEXDIGIT
	}
	for ch := 'a'; ch <= 'f'; ch++ {
		FLAGS[ch] |= IS_HEXDIGIT
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		FLAGS[ch] |= IS_ALPHA
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		FLAGS[ch] |= IS_ALPHA
	}

}

func (t *Tokenizer) Process() []Token {
	for t.pos < t.max {
		ch := t.charsToProcess[t.pos]
		if isAlphabetic(ch) {
			t.lexIdentifier()
		} else {
			st := string(ch)
			switch st {

			case "+":
				if t.isTwoCharToken(TokenKind{TokenKindType: INC, TokenChars: []rune(INC), HasPayload: len([]rune(INC)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: INC, TokenChars: []rune(INC), HasPayload: len([]rune(INC)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: INC, TokenChars: []rune(INC), HasPayload: len([]rune(INC)) == 0})
				}
				break
			case "-":
				if t.isTwoCharToken(TokenKind{TokenKindType: DEC, TokenChars: []rune(DEC), HasPayload: len([]rune(DEC)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: DEC, TokenChars: []rune(DEC), HasPayload: len([]rune(DEC)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: MINUS, TokenChars: []rune(MINUS), HasPayload: len([]rune(MINUS)) == 0})
				}
				break
			case ":":
				t.pushCharToken(TokenKind{TokenKindType: COLON, TokenChars: []rune(COLON), HasPayload: len([]rune(COLON)) == 0})
				break
			case ".":
				t.pushCharToken(TokenKind{TokenKindType: DOT, TokenChars: []rune(DOT), HasPayload: len([]rune(DOT)) == 0})
				break
			case ",":
				t.pushCharToken(TokenKind{TokenKindType: COMMA, TokenChars: []rune(COMMA), HasPayload: len([]rune(COMMA)) == 0})
				break
			case "*":
				t.pushCharToken(TokenKind{TokenKindType: STAR, TokenChars: []rune(STAR), HasPayload: len([]rune(STAR)) == 0})
				break
			case "/":
				t.pushCharToken(TokenKind{TokenKindType: DIV, TokenChars: []rune(DIV), HasPayload: len([]rune(DIV)) == 0})
				break
			case "%":
				t.pushCharToken(TokenKind{TokenKindType: MOD, TokenChars: []rune(MOD), HasPayload: len([]rune(MOD)) == 0})
				break
			case "(":
				t.pushCharToken(TokenKind{TokenKindType: LPAREN, TokenChars: []rune(LPAREN), HasPayload: len([]rune(LPAREN)) == 0})
				break
			case ")":
				t.pushCharToken(TokenKind{TokenKindType: RPAREN, TokenChars: []rune(RPAREN), HasPayload: len([]rune(RPAREN)) == 0})
				break
			case "[":
				t.pushCharToken(TokenKind{TokenKindType: LSQUARE, TokenChars: []rune(LSQUARE), HasPayload: len([]rune(LSQUARE)) == 0})
				break
			case "#":
				t.pushCharToken(TokenKind{TokenKindType: HASH, TokenChars: []rune(HASH), HasPayload: len([]rune(HASH)) == 0})
				break
			case "]":
				t.pushCharToken(TokenKind{TokenKindType: RSQUARE, TokenChars: []rune(RSQUARE), HasPayload: len([]rune(RSQUARE)) == 0})
				break
			case "{":
				//t.pushCharToken(TokenKind{TokenKindType: LCURLY, TokenChars: []rune(LCURLY), HasPayload: len([]rune(LCURLY)) == 0})
				//break
				t.pos++
				break
			case "}":
				//t.pushCharToken(TokenKind{TokenKindType: RCURLY, TokenChars: []rune(RCURLY), HasPayload: len([]rune(RCURLY)) == 0})
				//break
				t.pos++
				break
			case "@":
				t.pushCharToken(TokenKind{TokenKindType: BEAN_REF, TokenChars: []rune(BEAN_REF), HasPayload: len([]rune(BEAN_REF)) == 0})
				break
			case "!":
				if t.isTwoCharToken(TokenKind{TokenKindType: NE, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: NE, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0})
				} else if t.isTwoCharToken(TokenKind{TokenKindType: NE, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: PROJECT, TokenChars: []rune(NE), HasPayload: len([]rune(NE)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: NOT, TokenChars: []rune(NOT), HasPayload: len([]rune(NOT)) == 0})
				}
				break
			case "=":
				if t.isTwoCharToken(TokenKind{TokenKindType: EQ, TokenChars: []rune(EQ), HasPayload: len([]rune(EQ)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: EQ, TokenChars: []rune(EQ), HasPayload: len([]rune(EQ)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: ASSIGN, TokenChars: []rune(ASSIGN), HasPayload: len([]rune(ASSIGN)) == 0})
				}
				break
			case "&":
				if t.isTwoCharToken(TokenKind{TokenKindType: SYMBOLIC_AND, TokenChars: []rune(SYMBOLIC_AND), HasPayload: len([]rune(SYMBOLIC_AND)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: SYMBOLIC_AND, TokenChars: []rune(SYMBOLIC_AND), HasPayload: len([]rune(SYMBOLIC_AND)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: FACTORY_BEAN_REF, TokenChars: []rune(FACTORY_BEAN_REF), HasPayload: len([]rune(FACTORY_BEAN_REF)) == 0})
				}
				break
			case "|":
				if !t.isTwoCharToken(TokenKind{TokenKindType: SYMBOLIC_OR, TokenChars: []rune(SYMBOLIC_OR), HasPayload: len([]rune(SYMBOLIC_OR)) == 0}) {
					fmt.Errorf("")
				}
				t.pushPairToken(TokenKind{TokenKindType: SYMBOLIC_OR, TokenChars: []rune(SYMBOLIC_OR), HasPayload: len([]rune(SYMBOLIC_OR)) == 0})
				break
			case "$":
				//if t.isTwoCharToken(TokenKind{TokenKindType: SELECT_LAST, TokenChars: []rune(SELECT_LAST), HasPayload: len([]rune(SELECT_LAST)) == 0}) {
				//	t.pushPairToken(TokenKind{TokenKindType: SELECT_LAST, TokenChars: []rune(SELECT_LAST), HasPayload: len([]rune(SELECT_LAST)) == 0})
				//} else {
				//	t.lexIdentifier()
				//}
				//break
				//支持以# ，$开头
				t.pushCharToken(TokenKind{TokenKindType: HASH, TokenChars: []rune(HASH), HasPayload: len([]rune(HASH)) == 0})
				break
			case ">":
				if t.isTwoCharToken(TokenKind{TokenKindType: GE, TokenChars: []rune(GE), HasPayload: len([]rune(GE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: GE, TokenChars: []rune(GE), HasPayload: len([]rune(GE)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: GE, TokenChars: []rune(GE), HasPayload: len([]rune(GE)) == 0})
				}
				break
			case "<":
				if t.isTwoCharToken(TokenKind{TokenKindType: LE, TokenChars: []rune(LE), HasPayload: len([]rune(LE)) == 0}) {
					t.pushPairToken(TokenKind{TokenKindType: LE, TokenChars: []rune(LE), HasPayload: len([]rune(LE)) == 0})
				} else {
					t.pushCharToken(TokenKind{TokenKindType: LE, TokenChars: []rune(LE), HasPayload: len([]rune(LE)) == 0})
				}
				break
			case "0":
				t.lexNumericLiteral(ch == '0')
				break
			case "1":
				t.lexNumericLiteral(ch == '0')
				break
			case "2":
				t.lexNumericLiteral(ch == '0')
				break
			case "3":
				t.lexNumericLiteral(ch == '0')
				break
			case "4":
				t.lexNumericLiteral(ch == '0')
				break
			case "5":
				t.lexNumericLiteral(ch == '0')
				break
			case "6":
				t.lexNumericLiteral(ch == '0')
				break
			case "7":
				t.lexNumericLiteral(ch == '0')
				break
			case "8":
				t.lexNumericLiteral(ch == '0')
				break
			case "9":
				t.lexNumericLiteral(ch == '0')
				break
			case " ":
				t.pos++
				break
			case "\t":
				t.pos++
				break
			case "\r":
				t.pos++
				break
			case "\n":
				t.pos++
				break
			case "'":
				t.lexQuotedStringLiteral()
				break
			case "\"":
				t.lexDoubleQuotedStringLiteral()
				break
			case string(0):
				t.pos++
				break
			}

		}
	}
	return t.tokens
}

func (t *Tokenizer) isTwoCharToken(kind TokenKind) bool {
	return len(kind.TokenChars) == 2 && t.charsToProcess[t.pos] == kind.TokenChars[0] &&
		t.charsToProcess[t.pos+1] == kind.TokenChars[1]
}

func (t *Tokenizer) pushPairToken(kind TokenKind) {
	t.tokens = append(t.tokens, Token{Kind: kind, StartPos: t.pos, EndPos: t.pos + 2})
	t.pos += 2
}

func (t *Tokenizer) pushCharToken(kind TokenKind) {
	t.tokens = append(t.tokens, Token{Kind: kind, StartPos: t.pos, EndPos: t.pos + 1})
	t.pos++
}

func (t *Tokenizer) pushOneCharOrTwoCharToken(kind TokenKind, pos int, data []rune) {
	t.tokens = append(t.tokens, Token{Kind: kind, StartPos: pos, Data: string(data), EndPos: pos + len(kind.TokenKindType)})
}

func (t *Tokenizer) pushIntToken(data []rune, isLong bool, start int, end int) {
	if isLong {
		kind := TokenKind{TokenKindType: LITERAL_LONG, TokenChars: []rune(LITERAL_LONG), HasPayload: len([]rune(LITERAL_LONG)) == 0}
		t.tokens = append(t.tokens, Token{Kind: kind, StartPos: start, Data: string(data), EndPos: end})
	} else {
		kind := TokenKind{TokenKindType: LITERAL_INT, TokenChars: []rune(LITERAL_INT), HasPayload: len([]rune(LITERAL_INT)) == 0}
		t.tokens = append(t.tokens, Token{Kind: kind, StartPos: start, Data: string(data), EndPos: end})
	}

}

func isAlphabetic(ch rune) bool {
	char := int(ch)
	if char > 255 {
		return false
	}
	return (FLAGS[ch] & IS_ALPHA) != 0
}

//判断是否是数字
func isDigit(ch rune) bool {
	char := int(ch)
	if char > 255 {
		return false
	}
	return (FLAGS[ch] & IS_DIGIT) != 0
}

func isIdentifier(ch rune) bool {
	return isAlphabetic(ch) || isDigit(ch) || ch == '_' || ch == '$'
}

func isHexadecimalDigit(ch rune) bool {
	char := int(ch)
	if char > 255 {
		return false
	}
	return (FLAGS[ch] & IS_HEXDIGIT) != 0
}

func (t *Tokenizer) lexIdentifier() {
	start := t.pos
	t.pos++
	for t.pos < t.max && isIdentifier(t.charsToProcess[t.pos]) {
		t.pos++
	}
	runes := t.charsToProcess[start:t.pos]
	alternativeOperatorNames := []string{"DIV", "EQ", "GE", "GT", "LE", "LT", "MOD", "NE", "NOT"}
	if (t.pos-start) == 2 || (t.pos-start) == 3 {
		asString := strings.ToUpper(string(runes))
		idx := utils.BinarySearch(alternativeOperatorNames, asString)
		if idx >= 0 {
			//t.pushOneCharOrTwoCharToken(TokenKind.valueOf(asString), start, runes)
			return
		}
	}
	t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: IDENTIFIER, TokenChars: []rune(IDENTIFIER), HasPayload: len([]rune(IDENTIFIER)) == 0}, string(runes), start, t.pos})
}

//处理数字判断
func (t *Tokenizer) lexNumericLiteral(firstCharIsZero bool) {
	isReal := false
	start := t.pos
	if start == t.max-1 {
		t.pos++
		t.pushIntToken(t.subarray(start, t.pos), false, start, t.pos)
		return
	}
	ch := t.charsToProcess[t.pos+1]
	//处理16进制
	isHex := ch == 'x' || ch == 'X'
	if firstCharIsZero && isHex {
		t.pos = t.pos + 1
		t.pos++
		for isHexadecimalDigit(t.charsToProcess[t.pos]) && t.pos < t.max-1 {
			t.pos++
		}
		if t.isChar('L', 'l') {
			t.pushHexIntToken(t.subarray(start+2, t.pos), true, start, t.pos)
			t.pos++
		} else {
			t.pushHexIntToken(t.subarray(start+2, t.pos), false, start, t.pos)
		}
		return
	} else {
		t.pos++
		//迭代数字
		for isDigit(t.charsToProcess[t.pos]) && t.pos < t.max-1 {
			t.pos++
		}
		if t.pos == t.max-1 {
			t.pos++
			t.pushIntToken(t.subarray(start, t.pos), false, start, t.pos)
			return
		}
		ch = t.charsToProcess[t.pos]
		endOfNumber := t.pos
		if ch == '.' {
			isReal = true
			endOfNumber = t.pos
			t.pos++
			for isDigit(t.charsToProcess[t.pos]) && t.pos < t.max-1 {
				t.pos++
			}
			if t.pos == endOfNumber {
				t.pos = endOfNumber
				t.pushIntToken(t.subarray(start, t.pos), false, start, t.pos)
				return
			}
		}
		endOfNumber = t.pos
		if t.isChar('L', 'l') {
			if isReal {
				panic("Real number cannot be suffixed with a long (L or l) suffix")
			}
			t.pushIntToken(t.subarray(start, endOfNumber), true, start, endOfNumber)
			t.pos++
		} else if t.isExponentChar(t.charsToProcess[t.pos]) {
			isReal = true
			t.pos++
			possibleSign := t.charsToProcess[t.pos]
			if t.isSign(possibleSign) {
				t.pos++
			}
			t.pos++
			for isDigit(t.charsToProcess[t.pos]) && t.pos < t.max-1 {
				t.pos++
			}
			isFloat := false
			if t.isFloatSuffix(t.charsToProcess[t.pos]) {
				isFloat = true
				t.pos++
				endOfNumber = t.pos
			} else if t.isDoubleSuffix(t.charsToProcess[t.pos]) {
				t.pos++
				endOfNumber = t.pos
			}
			t.pushRealToken(t.subarray(start, t.pos), isFloat, start, t.pos)
		} else {
			ch := t.charsToProcess[t.pos]
			isFloat := false
			if t.isFloatSuffix(ch) {
				isReal = true
				isFloat = true
				t.pos++
				endOfNumber = t.pos
			} else if t.isDoubleSuffix(ch) {
				isReal = true
				t.pos++
				endOfNumber = t.pos
			}
			if t.pos == t.max-1 {
				t.pos++
				endOfNumber = t.pos
			}

			if isReal {
				t.pushRealToken(t.subarray(start, endOfNumber), isFloat, start, endOfNumber)
			} else {
				t.pushIntToken(t.subarray(start, endOfNumber), false, start, endOfNumber)
			}
		}
	}
}

func (t *Tokenizer) lexQuotedStringLiteral() {
	start := t.pos
	terminated := false
	for !terminated {
		t.pos++
		ch := t.charsToProcess[t.pos]
		if string(ch) == "'" {
			if t.pos < t.max-1 {
				if string(t.charsToProcess[t.pos+1]) == "'" {
					t.pos++
				} else {
					terminated = true
				}
			}
			terminated = true
			if t.pos == t.max {
				panic("Cannot find terminating '' for string")
			}
		}
	}
	t.pos++
	process := t.charsToProcess[start:t.pos]
	t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_STRING, TokenChars: []rune(LITERAL_STRING), HasPayload: len([]rune(LITERAL_STRING)) == 0}, string(process), start, t.pos})
}

func (t *Tokenizer) lexDoubleQuotedStringLiteral() {
	start := t.pos
	terminated := false
	for !terminated {
		t.pos++
		ch := t.charsToProcess[t.pos]
		if string(ch) == "\"" {
			if t.pos < t.max-1 {
				if string(t.charsToProcess[t.pos+1]) == "\"" {
					t.pos++
				} else {
					terminated = true
				}
			}
			terminated = true
			if t.pos == t.max {
				panic("Cannot find terminating \" for string")
			}
		}
	}
	t.pos++
	process := t.charsToProcess[start:t.pos]
	t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_STRING, TokenChars: []rune(LITERAL_STRING), HasPayload: len([]rune(LITERAL_STRING)) == 0}, string(process), start, t.pos})
}

func (t *Tokenizer) isExponentChar(ch rune) bool {
	return ch == 'e' || ch == 'E'
}

func (t *Tokenizer) isFloatSuffix(ch rune) bool {
	return ch == 'f' || ch == 'F'
}

func (t *Tokenizer) isDoubleSuffix(ch rune) bool {
	return ch == 'd' || ch == 'D'
}

func (t *Tokenizer) isSign(ch rune) bool {
	return ch == '+' || ch == '-'
}

func (t *Tokenizer) isChar(a rune, b rune) bool {
	r := t.charsToProcess[t.pos]
	return r == a || r == b

}

func (t *Tokenizer) pushHexIntToken(data []rune, isLong bool, start int, end int) {
	if len(data) == 0 {
		if isLong {
			panic("The value  cannot be parsed as a long")
		} else {
			panic("The value  cannot be parsed as a int")
		}
	}

	if isLong {
		t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_HEXLONG, TokenChars: []rune(LITERAL_HEXLONG), HasPayload: len([]rune(LITERAL_HEXLONG)) == 0}, string(data), start, end})
	} else {
		t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_HEXINT, TokenChars: []rune(LITERAL_HEXINT), HasPayload: len([]rune(LITERAL_HEXINT)) == 0}, string(data), start, end})
	}
}

func (t *Tokenizer) pushRealToken(data []rune, isFloat bool, start int, end int) {
	if isFloat {
		t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_REAL_FLOAT, TokenChars: []rune(LITERAL_REAL_FLOAT), HasPayload: len([]rune(LITERAL_REAL_FLOAT)) == 0}, string(data), start, end})
	} else {
		t.tokens = append(t.tokens, Token{TokenKind{TokenKindType: LITERAL_REAL, TokenChars: []rune(LITERAL_REAL), HasPayload: len([]rune(LITERAL_REAL)) == 0}, string(data), start, end})
	}
}

func (t *Tokenizer) subarray(start int, end int) []rune {
	result := make([]rune, 0)
	runes := t.charsToProcess[start:end]
	result = runes
	return result
}
