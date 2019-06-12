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

)
