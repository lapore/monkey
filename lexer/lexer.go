package lexer

import "fmt"
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

// This is a helpr print function...
func (l* Lexer) Print() {
    fmt.Printf("position = %d readPosition = %d ch = %c", l.position, l.readPosition, l.ch)
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

// Dont increase position | readPosition, but just give us what next char is!
func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    // Continuing to skip whitespace...
    for isWhitespace(l.ch) {
        l.readChar()
    }

    switch {
        case l.ch == '=':
            tok = newToken(token.ASSIGN, l.ch)
        case l.ch == '+':
            tok = newToken(token.PLUS, l.ch)
        case l.ch == '-':
            tok = newToken(token.MINUS, l.ch)
        case l.ch == '!':
            tok = newToken(token.BANG, l.ch)
        case l.ch == '*':
            tok = newToken(token.ASTERISK, l.ch)
        case l.ch == '/':
            tok = newToken(token.SLASH, l.ch)
        case l.ch == '<':
            tok = newToken(token.LT, l.ch)
        case l.ch == '>':
            tok = newToken(token.GT, l.ch)
        case l.ch == ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case l.ch == '(':
            tok = newToken(token.LPAREN, l.ch)
        case l.ch == ')':
            tok = newToken(token.RPAREN, l.ch)
        case l.ch == '{':
            tok = newToken(token.LBRACE, l.ch)
        case l.ch == '}':
            tok = newToken(token.RBRACE, l.ch)
        case l.ch == ',':
            tok = newToken(token.COMMA, l.ch)
        case l.ch == '+': {
            tok = newToken(token.PLUS, l.ch)
        }
        case l.ch == 0: {
            tok.Literal = ""
            tok.Type    = token.EOF
        }
        // The first character is a digit, aka [0-9]+
        case isDigit(l.ch) : {
            var leftposition int = l.position
            for isDigit(l.ch) {
                l.readChar()
            }
            tok.Literal = l.input[leftposition:l.position]
            tok.Type = token.INT

            // TODO: However if next letter is anything like "." or letter, it should be invalid??
            // We should report error...
            return tok
        }
        // The first character is a letter, aka [a-zA-Z_], the remaining can be [a-zA-Z_0-9]
        // either an [identifier | keyword]!
        case isLetter(l.ch): {
            var leftposition int = l.position
            for isLetter(l.ch) || isDigit(l.ch) {
                l.readChar()
            }
            tok.Literal = l.input[leftposition:l.position]
            tok.Type    = stringToType(tok.Literal)  // If they are keywords, return here
            return tok
        }
        default : {
            tok = newToken(token.ILLEGAL, l.ch)
        }
    }
    l.readChar()
    return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type : tokenType, Literal: string(ch)}
}

func stringToType(literal string) token.TokenType {
    //fmt.Printf("Calling stringTotype %q\n", literal)
    var keywords = map[string] token.TokenType {
        "fn"    : token.FUNCTION,
        "let"   : token.LET,
    }
    if tok, ok := keywords[literal]; ok {
        return tok
    }
    return token.IDENT
}

func isLetter(ch byte) bool {
    if ((ch >= 'a' && ch <= 'z')  ||
        (ch >= 'A' && ch <= 'Z')  ||
        (ch == '_'))  {
        return true
    }
    return false
}

func isDigit(ch byte) bool {
    if ch >= '0' && ch <= '9' {
        return true
    }
    return false
}

func isWhitespace(ch byte) bool {
   if ch == '\t' || ch == '\r' || ch == '\n' || ch == ' ' {
       return true
   }
   return false
}
