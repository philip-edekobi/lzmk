package main

import (
	"fmt"

	"github.com/philip-edekobi/lzmk/pkg/codegen"
	"github.com/philip-edekobi/lzmk/pkg/lexer"
	"github.com/philip-edekobi/lzmk/pkg/parser"
)

func main() {
	// fmt.Println("lzmk - Lazymark compiler (stub)")
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: lzmk [options] file...")
	// 	os.Exit(1)
	// }

	input := "# Sample Lazymark\n\n## Sample Title Heading(hehe)\n\nBody consists of \"Hello World!\"\n#! (url)[alternative text details]\n\n### author Philip\n### date 2025-09-04"

	fmt.Println(input)

	l := lexer.NewLexer(input)
	tokens, err := l.Lex()
	if err != nil {
		panic(err)
	}

	fmt.Println("TOKENS:\n", tokens)

	p := parser.NewParser(tokens)
	ast, err := p.Parse(tokens)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nAST:\n\n")
	ast.PrettyPrint()

	html, err := codegen.GenerateHTML(ast)
	if err != nil {
		panic(err)
	}

	fmt.Println("HTML:", html)
}
