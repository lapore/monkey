package repl

import (
    "bufio"
    "fmt"
    "io"
    "monkey/lexer"
    "monkey/token"
)

const (
    PROMPT = ">> "
)

//func Start(in io.Reader, out io.Writer) {
func Start(in io.Reader) {
    var scanner *bufio.Scanner
    scanner = bufio.NewScanner(in)

    // forever loop
    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if (!scanned) {
            return
        }

        line := scanner.Text()
        l := lexer.New(line)
        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}
