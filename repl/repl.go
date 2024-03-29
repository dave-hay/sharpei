package repl

import (
	"bufio"
	"fmt"
	"github.com/sharpei/object"
	"io"
	"log"

	"github.com/sharpei/evaluator"
	"github.com/sharpei/lexer"
	"github.com/sharpei/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if _, err := io.WriteString(out, evaluated.Inspect()); err != nil {
				log.Fatalln(err)
			}
			if _, err := io.WriteString(out, "\n"); err != nil {
				log.Fatalln(err)
			}
		}

	}
}

const DOG = `
  __    __
o-''))_____\\
"--__/ * * * )
c_c__/-c____/
`

const DOG2 = `
     |\_/|
     | @ @   Woof!
     |   <>              _
     |  _/\------____ ((| |))
     |               '--' |
____|_       ___|   |___.'
/_/_____/____/_______|
Look like Phil found an error!
 parser errors:
`

const ALIEN = `
                       .-.
        .-""'""-.    |(@ @)
	_/' oOoOoOoOo'\_ \ \-/
	'.-=-=-=-=-=-=-.' \/ \
      '-=.=-.-=.=-'    \ /\
         ^  ^  ^       _H_ \
`

const FLAME = `
		   _   (        __    (     ~->>
	,-----' |__,_~~___<'__')-~__--__-~->> <
	| //  : | -__   ~__ o)____)),__ - '> >-  >
	| //  : |- \_ \ -\_\ -\ \ \ ~\_  \ ->> - ,  >>
	| //  : |_~_\ -\__\ \~'\ \ \, \__ . -<-  >>
	'-----._| '  -__'-- - ~~ -- ' --~> >
	 _/___\_    //)_'//  | ||]
_____[_______]_[~~-_ (.L_/  ||
[____________________]' '\_,/'/
Look like You found an error!
 parser errors:
`

func printParserErrors(out io.Writer, errors []string) {
	//errMsg := DOG2 + "Look like Phil found an error!\n" + " parser errors:\n"
	//io.WriteString(out, DOG2)
	//io.WriteString(out, "Look like Phil found an error!\n")
	//io.WriteString(out, " parser errors:\n")
	io.WriteString(out, FLAME)
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
