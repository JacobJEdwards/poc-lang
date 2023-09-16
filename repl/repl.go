package repl 

import (
    "bufio"
    "fmt"
    "io"
    "github.com/jacobjedwards/poc-lang/lexer"
    "github.com/jacobjedwards/poc-lang/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)

    for {
        fmt.Print(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return 
        }

        line := scanner.Text()

        l := lexer.New(line)

        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Printf("%+v\n", tok)
        }
    }
}

func LexOutput(line string) {
    l := lexer.New(line)

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
        fmt.Printf("%+v\n", tok)
    }
}
