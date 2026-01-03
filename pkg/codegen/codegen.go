package codegen

import (
	"github.com/philip-edekobi/lzmk/pkg/parser"
)

func GenerateHTML(ast *parser.AST) (string, error) {
	page, err := generatePageString(ast)
	if err != nil {
		return "", err
	}

	return page, err
}

func GenerateReact(ast *parser.AST) (string, error) {
	return "", nil
}
