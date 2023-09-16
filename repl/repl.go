package repl 

import (
    "bufio"
    "fmt"
    "io"
    "github.com/jacobjedwards/poc-lang/lexer"
    "github.com/jacobjedwards/poc-lang/token"
    "github.com/jacobjedwards/poc-lang/parser"
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
        p := parser.New(l)

        program := p.ParseProgram()

        if len(p.Errors()) != 0 {
            printParserErrors(out, p.Errors())
            continue
        }

        io.WriteString(out, program.String())
        io.WriteString(out, "\n")
    }
}

func LexOutput(line string) {
    l := lexer.New(line)

    for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
        fmt.Printf("%+v\n", tok)
    }
}

const WARNING = `
                                                                                             
                                                                                        
                ░░░░                                                                    
                                                                                        
                                            ██                                          
                                          ██░░██                                        
  ░░          ░░                        ██░░░░░░██                            ░░░░      
                                      ██░░░░░░░░░░██                                    
                                      ██░░░░░░░░░░██                                    
                                    ██░░░░░░░░░░░░░░██                                  
                                  ██░░░░░░██████░░░░░░██                                
                                  ██░░░░░░██████░░░░░░██                                
                                ██░░░░░░░░██████░░░░░░░░██                              
                                ██░░░░░░░░██████░░░░░░░░██                              
                              ██░░░░░░░░░░██████░░░░░░░░░░██                            
                            ██░░░░░░░░░░░░██████░░░░░░░░░░░░██                          
                            ██░░░░░░░░░░░░██████░░░░░░░░░░░░██                          
                          ██░░░░░░░░░░░░░░██████░░░░░░░░░░░░░░██                        
                          ██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██                        
                        ██░░░░░░░░░░░░░░░░██████░░░░░░░░░░░░░░░░██                      
                        ██░░░░░░░░░░░░░░░░██████░░░░░░░░░░░░░░░░██                      
                      ██░░░░░░░░░░░░░░░░░░██████░░░░░░░░░░░░░░░░░░██                    
        ░░            ██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██                    
                        ██████████████████████████████████████████                      
                                                                                        
                                                                                        
                                                                                        
                                                                                        
                  ░░                                                                    
`

func printParserErrors(out io.Writer, errors []*parser.ParseError) {
    io.WriteString(out, WARNING)
    io.WriteString(out, "ERROR ERROR ERROR\n")
    io.WriteString(out, "parser errors:\n")
    for _, err := range errors {
        io.WriteString(out, "\t" + err.Error() + "\n")
    }
}
