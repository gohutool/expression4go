package expression4go

import (
	"container/list"
	"fmt"
	"strings"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : templateExpressionParser.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:06
* 修改历史 : 1. [2022/4/7 18:06] 创建文件 by NST
*/

type TemplateAwareExpressionParser struct {
	Bracket
}

type Bracket struct {
	bracket string
	pos     int
}

func (t *TemplateAwareExpressionParser) ParseExpression(expressionString string) Expression {
	return t.parseExpressionContext(expressionString, nil)
}

func (t *TemplateAwareExpressionParser) parseExpressionContext(expressionString string, context ParserContext) Expression {
	if context != nil && context.isTemplate() {
		return t.parseTemplate(expressionString, context)
	}
	return t.DoParseExpression(expressionString)
}

func (t *TemplateAwareExpressionParser) parseTemplate(expressionString string, context ParserContext) Expression {
	if expressionString == "" {
		return &LiteralExpression{}
	}
	expressions := t.parseExpressions(expressionString, context)
	if len(expressions) == 1 {
		return expressions[0]
	}
	return &CompositeStringExpression{ExpressionString: expressionString, Expressions: expressions}
}

func (t *TemplateAwareExpressionParser) parseExpressions(expressionString string, context ParserContext) []Expression {
	expressions := make([]Expression, 0)
	prefix := context.getExpressionPrefix()
	suffix := context.getExpressionSuffix()
	startIdx := 0
	if startIdx < len(expressionString) {
		prefixIndex := strings.Index(expressionString, prefix)
		if prefixIndex >= startIdx {
			if prefixIndex > startIdx {
				runes := []rune(expressionString)
				expressions = append(expressions, &LiteralExpression{literalValue: string(runes[startIdx:prefixIndex])})
			}
			afterPrefixIndex := prefixIndex + len(prefix)
			suffixIndex := skipToCorrectEndSuffix(suffix, expressionString, afterPrefixIndex)
			if suffixIndex == -1 {
				fmt.Errorf(expressionString, prefixIndex, "No ending suffix '"+
					suffix+"' for expression starting at character "+string(prefixIndex)+": "+string(expressionString[0:prefixIndex]))
			}

			if suffixIndex == afterPrefixIndex {
				fmt.Errorf(expressionString, prefixIndex, "No expression defined within delimiter '"+
					string(prefix)+string(suffix)+
					"' at character "+string(prefixIndex))
			}
			expr := expressionString[prefixIndex+len(prefix) : suffixIndex]
			if expr == "" {
				fmt.Errorf(expressionString, prefixIndex, "No expression defined within delimiter '"+
					string(prefix)+string(suffix)+
					"' at character "+string(prefixIndex))
			}
			expressions = append(expressions, t.ParseExpression(expressionString))
		}
	}
	return nil
}

func skipToCorrectEndSuffix(suffix string, expressionString string, afterPrefixIndex int) int {
	pos := afterPrefixIndex
	maxlen := len(expressionString)
	nextSuffix := strings.Index(expressionString, suffix)
	if nextSuffix == -1 {
		return -1
	} else {
		stack := list.List{}
		for pos < maxlen && (!isSuffixHere(expressionString, pos, suffix) || stack.Len() != 0) {
			pos++
			ch := string(expressionString[pos])
			switch ch {
			case "'":
			case "\\":
				{
					endLiteral := IndexOf(expressionString, ch, pos+1)
					if endLiteral == -1 {
						//err
					}
					pos = endLiteral
					break
				}
			case "(":
			case "[":
			case "{":
				{
					var bracket = Bracket{ch, pos}
					stack.PushFront(bracket)
					break
				}
			case ")":
			case "]":
			case "}":
				{
					if stack.Len() != 0 {
						fmt.Errorf(expressionString, pos, "Found closing"+ch+"'at postion'"+string(pos)+
							" without an opening"+theOpenBracketFor(ch)+"")
					}
					pop := stack.Front().Value
					if pop != nil {
						bracket := pop.(Bracket)
						closeBracket := bracket.compatibleWithCloseBracket(ch)
						if !closeBracket {

						}
					}
				}
			}

		}
		if stack.Len() != 0 {
			pop := stack.Front().Value
			fmt.Errorf(expressionString, pop.(Bracket).pos, "Missing closing"+theCloseBracketFor(pop.(Bracket).bracket)+"'for'"+
				pop.(Bracket).bracket+"'at postion'"+string(pop.(Bracket).pos))
		} else {
			if !isSuffixHere(expressionString, pos, suffix) {
				return -1
			}
			return pos
		}
	}
	return 0
}

func isSuffixHere(expressionString string, pos int, suffix string) bool {
	suffixPosition := 0
	for i := 0; i < len(suffix) && pos < len(expressionString); i++ {
		s := string(expressionString[pos])
		suffixPosition++
		s2 := string(suffix[suffixPosition])
		if s != s2 {
			return false
		}
	}
	return suffixPosition == len(suffix)
}

func (b *Bracket) compatibleWithCloseBracket(ch string) bool {
	if b.bracket == "{" {
		return ch == "}"
	}
	if b.bracket == "[" {
		return ch == "]"
	}
	if b.bracket == "{" {
		return ch == "}"
	} else {
		return ch == ")"
	}
}

func theOpenBracketFor(closeBracket string) string {
	if closeBracket == "}" {
		return "{"
	}
	if closeBracket == "]" {
		return "["
	}
	return "("
}

func theCloseBracketFor(openBracket string) string {
	if openBracket == "{" {
		return "{"
	}
	if openBracket == "[" {
		return "]"
	}
	return ")"
}
