package ast

import (
    "monkey/token"
)

type Node interface {
    // Type of Node in string
    TokenLiteral() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

// ------------------------------
// Expression: Identifier Node 
// ------------------------------
type Identifier struct {
    Token token.Token
    Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
    return i.Token.Literal
}

// ------------------------------
// Statement: Let Node 
// ------------------------------

// let a = 5+5;
type LetStatement struct {
    Token token.Token //let
    Name  *Identifier //a
    Value Expression  //5+5 
}

func (ls *LetStatement) statementNode() {

}
func (ls *LetStatement) TokenLiteral() string {
    return ls.Token.Literal
}

// ------------------------------
// Statement: return statement 
// ------------------------------

// return 5+5;
type ReturnStatement struct {
    Token       token.Token //return
    ReturnValue Expression  //5+5 
}

func (rs *ReturnStatement) statementNode() {

}
func (rs *ReturnStatement) TokenLiteral() string {
    return rs.Token.Literal
}



// A program consists of Multiple Statements node
type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}
