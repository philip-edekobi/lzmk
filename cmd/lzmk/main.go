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

	input := "# Sample Lazymark(TITLE)\n\n## Sample Section Heading(hehe)\n\nBody consists of \"Hello World!\"\n This is a naturally occuring body which is highlited by a lot of html elements. Most of these are all about tactics and wisdom as a developer. Make you no go dey gbezome bro. \n#! (url)[alternative text details]\n\n### author Philip\n### date 2025-09-04"

	fmt.Println(input)

	l := lexer.NewLexer(input)
	tokens, err := l.Lex()
	if err != nil {
		panic(err)
	}

	fmt.Println("TOKENS:\n", tokens)

	p := parser.NewParser(tokens)
	ast, err := p.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nAST:\n\n")
	ast.PrettyPrint()

	_, authorOk := p.MetaHashMap["author"]
	if !authorOk {
		panic("`author` missing in metadata")
	}

	_, dateOk := p.MetaHashMap["date"]
	if !dateOk {
		panic("`date` missing in metadata")
	}

	html, err := codegen.GenerateHTML(ast, p.MetaHashMap)
	if err != nil {
		panic(err)
	}

	fmt.Println("HTML:", html)
}
