package lexer

import (
    "testing"

    "monkey/token"
)

type CheckType struct {
    expectedType     token.TokenType
    expectedLiteral  string
}

func TestNextToken (t *testing.T) {
    var test_name = "TestNextToken"

    input := `let five = 5;
              let ten = 10;
              let add = fn(x, y) {
                  x + y;
              };
              let result = add(five, ten);`

    tests := []CheckType {
             {token.LET,          "let"},
             {token.IDENT,        "five"},
             {token.ASSIGN,       "="},
             {token.INT,          "5"},
             {token.SEMICOLON,    ";"},

             {token.LET,          "let"},
             {token.IDENT,        "ten"},
             {token.ASSIGN,       "="},
             {token.INT,          "10"},
             {token.SEMICOLON,    ";"},

             {token.LET,          "let"},
             {token.IDENT,        "add"},
             {token.ASSIGN,       "="},
             {token.FUNCTION,     "fn"},
             {token.LPAREN,       "("},
             {token.IDENT,        "x"},
             {token.COMMA,        ","},
             {token.IDENT,        "y"},
             {token.RPAREN,       ")"},
             {token.LBRACE,       "{"},

             {token.IDENT,        "x"},
             {token.PLUS,         "+"},
             {token.IDENT,        "y"},
             {token.SEMICOLON,    ";"},
             {token.RBRACE,       "}"},
             {token.SEMICOLON,    ";"},

             {token.LET,          "let"},
             {token.IDENT,        "result"},
             {token.ASSIGN,       "="},
             {token.IDENT,        "add"},
             {token.LPAREN,       "("},
             {token.IDENT,        "five"},
             {token.COMMA,        ","},
             {token.IDENT,        "ten"},
             {token.RPAREN,       ")"},
             {token.SEMICOLON,    ";"},
             {token.EOF,       ""},

           }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("%q[%d] - tokentype wrong. expected=%q, got=%q",
                     test_name, i, tt.expectedType, tok.Type)
        } else {
            t.Logf("%q[%d] - tokentype match. expected=%q, got=%q",
                     test_name, i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("%q[%d] - literal wrong. expected=%q, got=%q",
                     test_name, i, tt.expectedLiteral, tok.Literal)
        } else {
            t.Logf("%q[%d] - literal match. expected=%q, got=%q",
                     test_name, i, tt.expectedLiteral, tok.Literal)
        }
    }
}

func TestNextToken2 (t *testing.T) {
    var test_name = "TestNextToken2"

    input := `=+(){},;`  // Using `` means raw string literal

    tests := []CheckType {
             {token.ASSIGN,    "="},
             {token.PLUS,      "+"},
             {token.LPAREN,    "("},
             {token.RPAREN,    ")"},
             {token.LBRACE,    "{"},
             {token.RBRACE,    "}"},
             {token.COMMA,     ","},
             {token.SEMICOLON, ";"},
             {token.EOF,       ""},
           }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("%q[%d] - tokentype wrong. expected=%q, got=%q",
                     test_name, i, tt.expectedType, tok.Type)
        } else {
            t.Logf("%q[%d] - tokentype match. expected=%q, got=%q",
                     test_name, i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("%q[%d] - literal wrong. expected=%q, got=%q",
                     test_name, i, tt.expectedLiteral, tok.Literal)
        } else {
            t.Logf("%q[%d] - literal match. expected=%q, got=%q",
                     test_name, i, tt.expectedLiteral, tok.Literal)
        }
    }
}

func TestNextToken3 (t *testing.T) {
    var test_name = "TestNextToken3"

    input := `/+-!*/<>`  // Using `` means raw string literal

    tests := []CheckType {
             {token.SLASH,     "/"},
             {token.PLUS,      "+"},
             {token.MINUS,     "-"},
             {token.BANG,      "!"},
             {token.ASTERISK,  "*"},
             {token.SLASH,     "/"},
             {token.LT,        "<"},
             {token.GT,        ">"},
             {token.EOF,       ""},
           }

    l := New(input)

    for i, tt := range tests {
        tok := l.NextToken()

        if tok.Type != tt.expectedType {
            t.Fatalf("%q[%d] - tokentype wrong. expected=%q, got=%q",
                     test_name, i, tt.expectedType, tok.Type)
        } else {
            t.Logf("%q[%d] - tokentype match. expected=%q, got=%q",
                     test_name, i, tt.expectedType, tok.Type)
        }

        if tok.Literal != tt.expectedLiteral {
            t.Fatalf("%q[%d] - literal wrong. expected=%q, got=%q",
                     test_name, i, tt.expectedLiteral, tok.Literal)
        } else {
            t.Logf("%q[%d] - literal match. expected=%q, got=%q",
                     test_name, i, tt.expectedLiteral, tok.Literal)
        }
    }
}
