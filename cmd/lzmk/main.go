package main

import (
	"fmt"

	"github.com/philip-edekobi/lzmk/pkg/compiler"
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

	l := lexer.NewLexer(input)
	tokens, err := l.Lex()
	if err != nil {
		panic(err)
	}

	fmt.Println(tokens)

	ast := parser.Parse(tokens)
	html := compiler.CompileHTML(ast)

	fmt.Println(html)
}
