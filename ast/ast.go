//abstract syntax tree
package ast

import (
	"bytes"

	"github.com/MYKatz/PLZ/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var output bytes.Buffer
	for _, stmt := range p.Statements {
		output.WriteString(stmt.String())
	}
	return output.String()
}

//letstatement functions

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}
func (l *LetStatement) String() string {
	var output bytes.Buffer

	output.WriteString(l.TokenLiteral() + " ")
	output.WriteString(l.Name.String() + " ")
	output.WriteString("be ")
	if l.Value != nil {
		output.WriteString(l.Value.String())
	}
	output.WriteString(" plz")
	return output.String()
}

//returnstatement functions
func (r *ReturnStatement) statementNode() {}
func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}
func (r *ReturnStatement) String() string {
	var output bytes.Buffer

	output.WriteString(r.TokenLiteral() + " ")
	if r.Value != nil {
		output.WriteString(r.Value.String())
	}
	output.WriteString(" plz")
	return output.String()
}

//expressionstatement functions
func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	var output bytes.Buffer

	if es.Expression != nil {
		output.WriteString(es.Expression.String())
	}

	return output.String()
}

//identifier functions

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
