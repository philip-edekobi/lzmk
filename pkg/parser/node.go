package parser

type NodeType int8

const (
	_ NodeType = iota
	RootNode
	TitleNode
	BodyNode
	BodyChildNode
)

type Node struct {
	Kind        NodeType
	StringValue string
	Line        int
	Col         int
	Children    []*Node
}

func (n *Node) Value() string {
	return n.StringValue
}

func newNode(kind NodeType, val string, line, col int) *Node {
	return &Node{
		Kind:        kind,
		StringValue: val,
		Line:        line,
		Col:         col,
		Children:    []*Node{},
	}
}
