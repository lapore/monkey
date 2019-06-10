package lexer

import "monkey/token"

type Lexer struct {
    input         string
    position      int         // current position in input (current char)
    readPosition  int         // current reading position in input  (after current char)
    ch            byte        // current char under exam
}

func New(input string) *Lexer {
    l := &Lexer{input: input, position : 0, readPosition : 0, ch : 0}
    l.readChar()
    return l
}


// l *Lexer is a receiver argument..
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    // Note for go else should be in the same line as { because go automatically insert ; 
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    // Note readPosition is always faster by 1?
    l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
        case '=':
            tok = newToken(token.ASSIGN, l.ch)
        case ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case '(':
            tok = newToken(token.LPAREN, l.ch)
        case ')':
            tok = newToken(token.RPAREN, l.ch)
        case '{':
            tok = newToken(token.LBRACE, l.ch)
        case '}':
            tok = newToken(token.RBRACE, l.ch)
        case ',':
            tok = newToken(token.COMMA, l.ch)
        case '+':
            tok = newToken(token.PLUS, l.ch)
        case 0:
            tok.Literal = ""
            tok.Type    = token.EOF
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type : tokenType, Literal: string(ch)}
}

