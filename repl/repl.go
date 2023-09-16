package repl

import (
	"bufio"
	"fmt"
	"github.com/jacobjedwards/poc-lang/evaluator"
	"github.com/jacobjedwards/poc-lang/lexer"
	"github.com/jacobjedwards/poc-lang/parser"
	"github.com/jacobjedwards/poc-lang/token"
	"io"
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

		event := evaluator.Eval(program)
		if event != nil {
			io.WriteString(out, event.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func ParserOutput(p *parser.Parser, out io.Writer) {
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	io.WriteString(out, program.String())
	io.WriteString(out, "\n")
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
		io.WriteString(out, "\t"+err.Error()+"\n")
	}
}
