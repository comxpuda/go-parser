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
	NodeBody  *AstNode
	NodeValue string
}

func (p *Parser) Parse(s string) *AstNode {
	p.source = s
	p.tokenizer = &Tokenizer{
		source: s,
	}
	p.lookahead = p.tokenizer.getNextToken()
	return p.Program()
}

func (p *Parser) Program() *AstNode {
	if body, err := p.Literal(); err != nil {
		return nil
	} else {
		return &AstNode{
			NodeType: "Program",
			NodeBody: body,
		}
	}
}

func (p *Parser) StatementList() *AstNode {
	return nil
}

func (p *Parser) Statement() *AstNode {
	return nil
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
	return nil
}

func (p *Parser) BlockStatement() *AstNode {
	return nil
}

func (p *Parser) ExpressionStatement() *AstNode {
	return nil
}

func (p *Parser) Expression() *AstNode {
	return nil
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
	return nil
}

func (p *Parser) MultiplicativeExpression() *AstNode {
	return nil
}

func (p *Parser) _LogicalExpression() *AstNode {
	return nil
}
func (p *Parser) _BinaryExpression() *AstNode {
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
	return nil
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
	return nil
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

	p.lookahead = p.tokenizer.getNextToken()
	return token, nil
}
