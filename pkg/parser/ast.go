package parser

import "fmt"

type AST struct {
	Root *Node
}

func (ast *AST) PrettyPrint() {
	fmt.Printf("NodeType %v->%s\n", ast.Root.Kind, ast.Root.StringValue)

	for _, child := range ast.Root.Children {
		child.prettyPrint("")
	}
}

func newAST(rootNode *Node) *AST {
	return &AST{
		Root: rootNode,
	}
}
