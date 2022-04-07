package expression4go

import (
	"container/list"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : internalSpelExpressionParser.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 18:26
* 修改历史 : 1. [2022/4/7 18:26] 创建文件 by NST
*/

import (
	. "github.com/gohutool/expression4go/spel"
	. "github.com/gohutool/expression4go/spel/ast"
	. "github.com/gohutool/expression4go/spel/standard"
)

type InternalSpelExpressionParser struct {
	*TemplateAwareExpressionParser

	expressionString string

	tokenStreamLength int

	tokenStreamPointer int

	Configuration SpelParserConfiguration

	tokenStream []Token

	constructedNodes list.List
}

func (i *InternalSpelExpressionParser) DoParseExpression(expressionString string) Expression {

	i.expressionString = expressionString
	tokenizer := Tokenizer{ExpressionString: expressionString}
	tokenizer.InitTokenizer()
	i.tokenStream = tokenizer.Process()
	i.tokenStreamLength = len(i.tokenStream)
	i.tokenStreamPointer = 0
	i.constructedNodes.Init()
	expression, err := i.eatExpression()
	if err != nil {
		panic("No node")
	}
	spelExpression := SpelExpression{Expression: expressionString, Ast: expression, Configuration: i.Configuration}
	return &spelExpression
}

func (i *InternalSpelExpressionParser) takeToken() Token {
	if i.tokenStreamPointer >= i.tokenStreamLength {
		panic("No token")
	}
	token := i.tokenStream[i.tokenStreamPointer]
	i.tokenStreamPointer++
	return token
}

func (i *InternalSpelExpressionParser) nextToken() (Token, error) {
	if i.tokenStreamPointer >= i.tokenStreamLength {
		return Token{}, ExpressionErr{}
	}
	token := i.tokenStream[i.tokenStreamPointer]
	i.tokenStreamPointer++
	return token, nil
}

func (i *InternalSpelExpressionParser) peekToken() (Token, error) {
	if i.tokenStreamPointer >= i.tokenStreamLength {
		return Token{}, ExpressionErr{}
	}
	token := i.tokenStream[i.tokenStreamPointer]
	return token, nil
}

func (i *InternalSpelExpressionParser) maybeEatRelationalOperator() (Token, error) {
	token, err := i.peekToken()
	//BUG修复 token不为空时处理
	if err == nil {
		if token.IsNumericRelationalOperator() {
			return token, nil
		}
	}
	return token, ExpressionErr{}
}

func (i *InternalSpelExpressionParser) eatRelationalExpression() (SpelNode, error) {
	expr, _ := i.eatSumExpression()
	relationalOperatorToken, err1 := i.maybeEatRelationalOperator()
	if err1 == nil {
		t := i.takeToken()
		rhExpr, _ := i.eatSumExpression()
		checkOperands(t, expr, rhExpr)
		kindType := relationalOperatorToken.Kind.TokenKindType

		if relationalOperatorToken.IsNumericRelationalOperator() {
			pos := toPos(t.StartPos, t.EndPos)
			nodes := make([]SpelNode, 0)
			nodes = append(nodes, expr)
			nodes = append(nodes, rhExpr)
			spelNodeImpl := SpelNodeImpl{Children: nodes}
			spelNodeImpl.Pos = pos
			operator := Operator{SpelNodeImpl: &spelNodeImpl}
			//不等
			if kindType == NE {
				eq := OpNE{Operator: &operator}
				//eq.Parent = &eq
				return &eq, nil
			}
			//相等
			if kindType == EQ {
				eq := OpEQ{Operator: &operator}
				//eq.Parent = &eq
				return &eq, nil
			}
			//大于
			if kindType == GT {
				eq := OpGE{Operator: &operator}
				//eq.Parent = &eq
				return &eq, nil
			}
			//小于
			if kindType == LT {
				eq := OpLE{Operator: &operator}
				//eq.Parent = &eq
				return &eq, nil
			}
			//大于等于
			if kindType == GE {
				eq := OpGE{Operator: &operator}
				//eq.Parent = &eq
				return &eq, nil
			}
			//小于等于
			if kindType == LE {
				eq := OpLE{Operator: &operator}
				//eq.Parent = &eq
				return &eq, nil
			}
		}
	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatSumExpression() (SpelNode, error) {
	expr, err := i.eatProductExpression()
	if err != nil && i.peekTokenTwo(INC, DEC) {

	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatProductExpression() (SpelNode, error) {
	expr, err := i.eatPowerIncDecExpression()
	if err != nil && i.peekTokenTwo(INC, DEC) {

	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatPowerIncDecExpression() (SpelNode, error) {
	expr, err := i.eatUnaryExpression()
	if err != nil && i.peekTokenTwo(INC, DEC) {

	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatUnaryExpression() (SpelNode, error) {
	if i.peekTokens(PLUS, MINUS, NOT) {
		t := i.takeToken()
		_, err := i.eatUnaryExpression()
		if err != nil {
			panic("No node")
		}
		if t.Kind.TokenKindType == NOT {
		}
	}
	return i.eatPrimaryExpression()
}

func (i *InternalSpelExpressionParser) eatPrimaryExpression() (SpelNode, error) {
	start, err := i.eatStartNode()
	node, err := i.eatNode()
	var nodes []SpelNode
	for node != nil {
		if nodes == nil {
			nodes = append(nodes, start)

		}
		nodes = append(nodes, node)
		node, _ = i.eatNode()
	}
	if start == nil || len(nodes) == 0 {
		return start, err
	}
	impl := SpelNodeImpl{Pos: toPos(start.GetStartPosition(), nodes[len(nodes)-1].GetEndPosition()), Children: nodes}
	expression := CompoundExpression{&impl}
	return &expression, nil
}

func (i *InternalSpelExpressionParser) eatStartNode() (SpelNode, error) {

	if i.maybeEatLiteral() {
		return i.pop(), nil
	} else if i.maybeEatFunctionOrVar() {
		return i.pop(), nil
	} else if i.maybeEatIndexer() {
		return i.pop(), nil
	}
	return nil, nil
}
func (i *InternalSpelExpressionParser) maybeEatLiteral() bool {
	t, err := i.peekToken()
	if err != nil {
		return false
	}
	kindType := t.Kind.TokenKindType
	if kindType == LITERAL_LONG {
		value, err := strconv.ParseInt(t.Data, 10, 32)
		if err != nil {
			// 将 int64 转化为 int
			value := *(*int)(unsafe.Pointer(&value))
			pos := toPos(t.StartPos, t.EndPos)
			spelNodeImpl := SpelNodeImpl{Pos: pos}
			typedValue := TypedValue{Value: value}
			l := Literal{OriginalValue: t.Data, SpelNodeImpl: &spelNodeImpl, Value: typedValue}
			literal := IntLiteral{Literal: &l}
			i.push(literal)
		}
	} else if kindType == LITERAL_INT {
		value, err := strconv.ParseInt(t.Data, 10, 64)
		if err == nil {
			// 将 int64 转化为 int
			value := *(*int)(unsafe.Pointer(&value))
			pos := toPos(t.StartPos, t.EndPos)
			spelNodeImpl := SpelNodeImpl{Pos: pos}
			typedValue := TypedValue{Value: value}
			l := Literal{OriginalValue: t.Data, SpelNodeImpl: &spelNodeImpl, Value: typedValue}
			literal := IntLiteral{Literal: &l}
			i.push(literal)
		}
	} else if kindType == LITERAL_REAL {
		value, err := strconv.ParseFloat(t.Data, 64)
		if err == nil {
			pos := toPos(t.StartPos, t.EndPos)
			spelNodeImpl := SpelNodeImpl{Pos: pos}
			typedValue := TypedValue{Value: value}
			l := Literal{OriginalValue: t.Data, SpelNodeImpl: &spelNodeImpl, Value: typedValue}
			literal := FloatLiteral{Literal: &l}
			i.push(literal)
		}
	} else if kindType == LITERAL_REAL_FLOAT {
		value, err := strconv.ParseFloat(t.Data, 64)
		if err == nil {
			pos := toPos(t.StartPos, t.EndPos)
			spelNodeImpl := SpelNodeImpl{Pos: pos}
			typedValue := TypedValue{Value: value}
			l := Literal{OriginalValue: t.Data, SpelNodeImpl: &spelNodeImpl, Value: typedValue}
			literal := FloatLiteral{Literal: &l}
			i.push(literal)
		}
	} else if kindType == LITERAL_STRING {
		data := t.Data
		valueWithinQuotes := data[1 : len(data)-1]
		valueWithinQuotes = strings.ReplaceAll(valueWithinQuotes, "''", "'")
		valueWithinQuotes = strings.ReplaceAll(valueWithinQuotes, "\"\"", "\"")
		typedValue := TypedValue{Value: valueWithinQuotes}
		pos := toPos(t.StartPos, t.EndPos)
		spelNodeImpl := SpelNodeImpl{Pos: pos}
		l := Literal{OriginalValue: valueWithinQuotes, SpelNodeImpl: &spelNodeImpl, Value: typedValue}
		literal := StringLiteral{Literal: &l}
		i.push(literal)
	} else {
		return false
	}
	i.nextToken()
	return true
}

func (i *InternalSpelExpressionParser) maybeEatFunctionOrVar() bool {
	if !i.peekTokenOnly(HASH) {
		return false
	}
	token := i.takeToken()

	functionOrVariableName := i.eatToken(IDENTIFIER)
	args := i.maybeEatMethodArgs()
	if args == nil {
		reference := VariableReference{Name: functionOrVariableName.StringValue()}
		pos := toPos(token.StartPos, functionOrVariableName.EndPos)
		impl := SpelNodeImpl{Pos: pos}
		reference.SpelNodeImpl = &impl
		i.push(reference)
		return true
	}
	return true
}

func (i *InternalSpelExpressionParser) maybeEatMethodOrProperty(nullSafeNavigation bool) bool {
	if i.peekTokenOnly(IDENTIFIER) {
		methodOrPropertyName := i.takeToken()
		impl := SpelNodeImpl{Pos: toPos(methodOrPropertyName.StartPos, methodOrPropertyName.EndPos)}
		i.push(PropertyOrFieldReference{NullSafe: nullSafeNavigation, Name: methodOrPropertyName.Data, SpelNodeImpl: &impl})
		return true
	}
	return false
}

func (i *InternalSpelExpressionParser) maybeEatIndexer() bool {

	//bug修复 获取token提前，否则报越界
	token, err := i.peekToken()
	if !i.peekTokenMatched(LSQUARE, true) {
		return false
	}
	equal := reflect.DeepEqual(token, Token{})

	if equal {
		panic("No token")
	}
	expr, err := i.eatExpression()
	if err != nil {
		panic("No node")
	}
	i.eatToken(RSQUARE)
	impl := SpelNodeImpl{Pos: toPos(token.StartPos, token.EndPos), Children: []SpelNode{expr}}
	i.push(Indexer{SpelNodeImpl: &impl})
	return true
}

func (i *InternalSpelExpressionParser) eatNode() (SpelNode, error) {
	two := i.peekTokenTwo(DOT, SAFE_NAVI)
	if two {
		return i.eatDottedNode(), nil
	}
	return i.eatNonDottedNode(), nil
}

//包含"."
func (i *InternalSpelExpressionParser) eatDottedNode() SpelNode {
	t := i.takeToken()
	nullSafeNavigation := t.Kind.TokenKindType == SAFE_NAVI
	if i.maybeEatMethodOrProperty(nullSafeNavigation) || i.maybeEatFunctionOrVar() {
		return i.pop()
	}
	if _, err := i.peekToken(); err != nil {
		panic("Unexpectedly ran out of input")
	} else {
		panic("Unexpected data after ''.'': ''{0}''")
	}
}

func (i *InternalSpelExpressionParser) eatNonDottedNode() SpelNode {
	if i.peekTokenOnly(LSQUARE) {
		if i.maybeEatIndexer() {
			return i.pop()
		}
	}
	return nil
}

func toPos(start int, end int) int {
	return (start << 16) + end
}

func (i *InternalSpelExpressionParser) push(newNode SpelNode) {
	i.constructedNodes.PushFront(newNode)
}

func (i *InternalSpelExpressionParser) pop() SpelNode {
	return i.constructedNodes.Front().Value.(SpelNode)
}

func (i *InternalSpelExpressionParser) maybeEatMethodArgs() []SpelNodeImpl {
	if !i.peekTokenOnly(LPAREN) {
		return nil
	}
	//args := make([]SpelNodeImpl,0)
	return nil
}

func (i *InternalSpelExpressionParser) consumeArguments(accumulatedArguments []SpelNodeImpl) {
	token, err := i.peekToken()
	if err != nil {
		panic("Expected token")
	}
	//var next Token
	i.nextToken()
	token, err = i.peekToken()
	if err == nil {
		panic("Unexpectedly ran out of arguments")
	}
	if token.Kind.TokenKindType != RPAREN {
		accumulatedArguments = append(accumulatedArguments)
	}

}
func (i *InternalSpelExpressionParser) eatExpression() (SpelNode, error) {
	expr, _ := i.eatLogicalOrExpression()
	_, err := i.peekToken()
	if err == nil {

	}
	//bug修复
	if expr == nil {
		return nil, nil
	}
	return expr, nil
}

func (i *InternalSpelExpressionParser) eatLogicalOrExpression() (SpelNode, error) {
	expr, _ := i.eatLogicalAndExpression()
	var result SpelNode
	result = expr
	for i.peekIdentifierToken("or") || i.peekTokenOnly(SYMBOLIC_OR) {
		t := i.takeToken()
		rhExpr, _ := i.eatRelationalExpression()
		checkOperands(t, expr, rhExpr)
		pos := toPos(t.StartPos, t.EndPos)
		nodes := make([]SpelNode, 0)
		nodes = append(nodes, expr)
		nodes = append(nodes, rhExpr)
		spelNodeImpl := SpelNodeImpl{Children: nodes}
		spelNodeImpl.Pos = pos
		operator := Operator{SpelNodeImpl: &spelNodeImpl}
		expr := OpOr{&operator}
		result = &expr
	}
	return result, nil
}

func (i *InternalSpelExpressionParser) eatLogicalAndExpression() (SpelNode, error) {
	expr, _ := i.eatRelationalExpression()
	var result SpelNode
	result = expr
	for i.peekIdentifierToken("and") || i.peekTokenOnly(SYMBOLIC_AND) {
		t := i.takeToken()
		rhExpr, _ := i.eatRelationalExpression()
		checkOperands(t, expr, rhExpr)
		pos := toPos(t.StartPos, t.EndPos)
		nodes := make([]SpelNode, 0)
		nodes = append(nodes, expr)
		nodes = append(nodes, rhExpr)
		spelNodeImpl := SpelNodeImpl{Children: nodes}
		spelNodeImpl.Pos = pos
		operator := Operator{SpelNodeImpl: &spelNodeImpl}
		expr := OpAnd{&operator}
		result = &expr
	}
	return result, nil
}

func (i *InternalSpelExpressionParser) eatToken(expectedKind TokenKindType) Token {
	token, err := i.nextToken()
	if err != nil {
		panic("Unexpectedly ran out of input")
	}
	if token.Kind.TokenKindType != expectedKind {
		panic("Unexpected token.")
	}
	return token
}
func (i *InternalSpelExpressionParser) peekTokenOnly(possible1 TokenKindType) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	return token.Kind.TokenKindType == possible1
}

func (i *InternalSpelExpressionParser) peekTokenMatched(desiredTokenKind TokenKindType, consumeIfMatched bool) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	if token.Kind.TokenKindType == desiredTokenKind {
		if consumeIfMatched {
			i.tokenStreamPointer++
		}
		return true
	}
	if desiredTokenKind == IDENTIFIER {
		return true

	}
	return false
}

func (i *InternalSpelExpressionParser) peekTokenTwo(possible1 TokenKindType, possible2 TokenKindType) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	return (token.Kind.TokenKindType == possible1) || token.Kind.TokenKindType == possible2
}
func (i *InternalSpelExpressionParser) peekTokens(possible1 TokenKindType, possible2 TokenKindType, possible3 TokenKindType) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	return (token.Kind.TokenKindType == possible1) || token.Kind.TokenKindType == possible2 || token.Kind.TokenKindType == possible3
}

func (i *InternalSpelExpressionParser) peekIdentifierToken(identifierString string) bool {
	token, err := i.peekToken()
	if err != nil {
		return false
	}
	return token.Kind.TokenKindType == IDENTIFIER && token.Data == identifierString
}

func checkOperands(token Token, left SpelNode, right SpelNode) {
	checkLeftOperand(token, left)
	checkRightOperand(token, right)
}

func checkLeftOperand(token Token, left SpelNode) {
	if left == nil {
		panic("Problem parsing left operand")
	}
}

func checkRightOperand(token Token, right SpelNode) {
	if right == nil {
		panic("Problem parsing right operand")
	}
}
