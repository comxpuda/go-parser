package main

import (
	"fmt"
)

type Parser struct {
	source    string
	tokenizer *Tokenizer
	lookahead *Token
}

type AstNode struct {
	NodeType  string
	NodeBody  []*AstNode
	NodeValue string
	// binaryexpression
	NodeOperator string
	NodeLeft     *AstNode
	NodeRight    *AstNode
}

func (p *Parser) Parse(s string) *AstNode {
	p.source = s
	p.tokenizer = &Tokenizer{
		source: s,
	}
	p.lookahead, _ = p.tokenizer.getNextToken()
	return p.Program()
}

func (p *Parser) Program() *AstNode {
	return &AstNode{
		NodeType: "Program",
		NodeBody: p.StatementList("NULL"),
	}
}

func (p *Parser) StatementList(stopLookahead string) []*AstNode {
	statementList := []*AstNode{p.Statement()}
	for p.lookahead != nil && p.lookahead.TokenType != stopLookahead {
		statementList = append(statementList, p.Statement())
	}
	return statementList
}

func (p *Parser) Statement() *AstNode {
	switch p.lookahead.TokenType {
	case ";":
		return p.EmptyStatement()
	case "{":
		return p.BlockStatement()
	default:
		return p.ExpressionStatement()
	}
}

func (p *Parser) ClassDeclaration() *AstNode {
	return nil
}

func (p *Parser) ClassExtends() *AstNode {
	return nil
}
func (p *Parser) FunctionDeclaration() *AstNode {
	return nil
}
func (p *Parser) FormalParameterList() *AstNode {
	return nil
}
func (p *Parser) ReturnStatement() *AstNode {
	return nil
}
func (p *Parser) IterationStatement() *AstNode {
	return nil
}
func (p *Parser) WhileStatement() *AstNode {
	return nil
}
func (p *Parser) DoWhileStatement() *AstNode {
	return nil
}
func (p *Parser) ForStatement() *AstNode {
	return nil
}

func (p *Parser) ForStatementInit() *AstNode {
	return nil
}

func (p *Parser) BreakStatement() *AstNode {
	return nil
}

func (p *Parser) ContinueStatement() *AstNode {
	return nil
}

func (p *Parser) IfStatement() *AstNode {
	return nil
}

func (p *Parser) VariableStatementInit() *AstNode {
	return nil
}

func (p *Parser) VariableStatement() *AstNode {
	return nil
}

func (p *Parser) VariableDeclarationList() *AstNode {
	return nil
}

func (p *Parser) VariableDeclaration() *AstNode {
	return nil
}

func (p *Parser) VariableInitializer() *AstNode {
	return nil
}

func (p *Parser) EmptyStatement() *AstNode {
	p.Eat(";")
	return &AstNode{
		NodeType: "EmptyStatement",
	}
}

func (p *Parser) BlockStatement() *AstNode {
	p.Eat("{")
	var body []*AstNode
	if p.lookahead.TokenType != "}" {
		body = p.StatementList("}")
	} else {
		body = []*AstNode{}
	}
	p.Eat("}")
	return &AstNode{
		NodeType: "BlockStatement",
		NodeBody: body,
	}
}

func (p *Parser) ExpressionStatement() *AstNode {
	exp := p.Expression()
	p.Eat(";")
	return &AstNode{
		NodeType: "ExpressionStatement",
		NodeBody: []*AstNode{exp},
	}
}

func (p *Parser) Expression() *AstNode {
	return p.AdditiveExpression()
}

func (p *Parser) AssignmentExpression() *AstNode {
	return nil
}

func (p *Parser) Identifier() *AstNode {
	return nil
}

func (p *Parser) _checkValidAssignmentTarget() *AstNode {
	return nil
}

func (p *Parser) _isAssignmentOperator() *AstNode {
	return nil
}

func (p *Parser) AssignmentOperator() *AstNode {
	return nil
}

func (p *Parser) LogicalORExpression() *AstNode {
	return nil
}

func (p *Parser) LogicalANDExpression() *AstNode {
	return nil
}

func (p *Parser) EqualityExpression() *AstNode {
	return nil
}

func (p *Parser) RelationalExpression() *AstNode {
	return nil
}

func (p *Parser) AdditiveExpression() *AstNode {
	left := p.MultiplicativeExpression()
	for p.lookahead.TokenType == "ADDITIVE_OPERATOR" {
		operator, _ := p.Eat("ADDITIVE_OPERATOR")
		right := p.MultiplicativeExpression()

		left = &AstNode{
			NodeType:     "BinaryExpression",
			NodeOperator: operator.TokenValue,
			NodeLeft:     left,
			NodeRight:    right,
		}
	}
	return left
}

func (p *Parser) MultiplicativeExpression() *AstNode {
	left := p.PrimaryExpression()
	for p.lookahead.TokenType == "MULTIPLICATIVE_OPERATOR" {
		operator, _ := p.Eat("MULTIPLICATIVE_OPERATOR")
		right := p.PrimaryExpression()

		left = &AstNode{
			NodeType:     "BinaryExpression",
			NodeOperator: operator.TokenValue,
			NodeLeft:     left,
			NodeRight:    right,
		}
	}
	return left
}

func (p *Parser) _LogicalExpression() *AstNode {
	return nil
}

func (p *Parser) UnaryExpression() *AstNode {
	return nil
}
func (p *Parser) LeftHandSideExpression() *AstNode {
	return nil
}
func (p *Parser) CallMemberExpression() *AstNode {
	return nil
}

func (p *Parser) _CallExpression() *AstNode {
	return nil
}
func (p *Parser) Arguments() *AstNode {
	return nil
}
func (p *Parser) ArgumentList() *AstNode {
	return nil
}
func (p *Parser) MemberExpression() *AstNode {
	return nil
}
func (p *Parser) PrimaryExpression() *AstNode {
	switch p.lookahead.TokenType {
	case "(":
		return p.ParenthesizedExpression()
	default:
		if node, err := p.Literal(); err == nil {
			return node
		}
		return nil
	}
}

func (p *Parser) NewExpression() *AstNode {
	return nil
}
func (p *Parser) ThisExpression() *AstNode {
	return nil
}
func (p *Parser) Super() *AstNode {
	return nil
}
func (p *Parser) _isLiteral() *AstNode {
	return nil
}
func (p *Parser) ParenthesizedExpression() *AstNode {
	p.Eat("(")
	exp := p.Expression()
	p.Eat(")")
	return exp
}

func (p *Parser) Literal() (*AstNode, error) {
	switch p.lookahead.TokenType {
	case "NUMBER":
		return p.NumericLiteral(), nil
	case "STRING":
		return p.StringLiteral(), nil
	}
	return nil, fmt.Errorf("Literal: unexpected literal production")
}

func (p *Parser) BooleanLiteral() *AstNode {
	return nil
}

func (p *Parser) NullLiteral() *AstNode {
	return nil
}

func (p *Parser) StringLiteral() *AstNode {
	if token, err := p.Eat("STRING"); err != nil {
		return nil
	} else {
		return &AstNode{
			NodeType:  "STRING",
			NodeValue: token.TokenValue[1 : len(token.TokenValue)-1],
		}
	}
}

func (p *Parser) NumericLiteral() *AstNode {
	if token, err := p.Eat("NUMBER"); err != nil {
		return nil
	} else {
		return &AstNode{
			NodeType:  "NUMBER",
			NodeValue: token.TokenValue,
		}
	}
}

func (p *Parser) Eat(tokenType string) (*Token, error) {
	token := p.lookahead

	if token == nil {
		return nil, fmt.Errorf("unexpected end of input,expected: %s", tokenType)
	}

	if token.TokenType != tokenType {
		return nil, fmt.Errorf("unexpected token: %s, expected is: %s", token.TokenValue, tokenType)
	}

	p.lookahead, _ = p.tokenizer.getNextToken()
	return token, nil
}
