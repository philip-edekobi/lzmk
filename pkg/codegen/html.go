package codegen

import (
	"fmt"

	"github.com/philip-edekobi/lzmk/pkg/parser"
)

func generatePageString(ast *parser.AST, metaHashMap map[string]string) (string, error) {
	page := ""
	author := metaHashMap["author"]
	date := metaHashMap["date"]

	root := ast.Root
	if len(root.Children) != 2 {
		return page, fmt.Errorf("Problem with parsing: root does not have 2 children")
	}

	title := root.Children[0].StringValue

	page, err := initializePage(
		page,
		title,
		author,
		date,
	)
	if err != nil {
		return "", err
	}

	for _, node := range root.Children[1].Children {
		if node.Kind == parser.MetadataNode {
			continue
		}

		genCode, err := generateHyperTextForNodeType(node)
		if err != nil {
			return page, err
		}

		page += genCode
	}

	page = closePage(page,
		title,
		author,
		date,
	)

	return page, nil
}

func initializePage(page, title, author, date string) (string, error) {
	if len(page) != 0 {
		return "", fmt.Errorf("Prior info already exists: possible page already initialized")
	}

	return page + generateInitalBoilerPlate(title, author, date), nil
}

func closePage(page, title, author, date string) string {
	return page + generateClosingBoilerPlate(author, date)
}
