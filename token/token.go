package token

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL   = "ILLEGAL"
    EOF       = "EOF"

    // Identifiers + literals
    IDENT     = "IDENT"
    INT       = "INT"

    ASSIGN    = "="
    PLUS      = "+"
    MINUS     = "-"
    BANG      = "!"
    ASTERISK  = "*"
    SLASH     = "/"

    LT        = "<"
    GT        = ">"
    L_EQ      = "<=" // TODO not implemented
    G_EQ      = ">=" // TODO not implemented
    EQ        = "=="
    NOT_EQ    = "!="

    COMMA     = ","
    SEMICOLON = ";"

    LPAREN    = "("
    RPAREN    = ")"
    LBRACE    = "{"
    RBRACE    = "}"

    FUNCTION  = "FUNCTION"
    LET       = "LET"
    IF        = "IF"
    ELSE      = "ELSE"
    RETURN    = "RETURN"
    TRUE      = "TRUE"
    FALSE     = "FALSE"

)
