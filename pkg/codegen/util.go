package codegen

import (
	"fmt"

	"github.com/philip-edekobi/lzmk/pkg/parser"
)

func generateHyperTextForNodeType(node *parser.Node) (string, error) {
	switch node.Kind {
	case parser.HeadingNode:
		return "", nil
	default:
		return "", fmt.Errorf("found unidentified node type")
	}
}
