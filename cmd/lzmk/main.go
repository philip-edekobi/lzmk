package main

import (
	"fmt"
	"os"

	"github.com/philip-edekobi/lzmk/pkg/compiler"
	"github.com/philip-edekobi/lzmk/pkg/lexer"
	"github.com/philip-edekobi/lzmk/pkg/parser"
)

func main() {
	fmt.Println("lzmk - Lazymark compiler (stub)")
	if len(os.Args) < 2 {
		fmt.Println("Usage: lzmk [options] file...")
		os.Exit(1)
	}

	input := "# Sample Lazymark\n\nHello World!"
	tokens := lexer.Lex(input)
	ast := parser.Parse(tokens)
	html := compiler.CompileHTML(ast)

	fmt.Println(html)
}
