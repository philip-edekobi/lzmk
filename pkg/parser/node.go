package parser

import "fmt"

type NodeType int8

const (
	_ NodeType = iota
	RootNode
	TitleNode
	BodyNode

	// BodyNode children

	HeadingNode
	TextNode
	MetadataNode
	URLNode
)

type URLInfo struct {
	AltText string
	URL     string
}

type MetadataInfo struct {
	Key   string
	Value string
}

type Node struct {
	Kind        NodeType
	StringValue string
	URLData     URLInfo
	Metadata    MetadataInfo
	Children    []*Node
}

func (n *Node) Value() string {
	return n.StringValue
}

func (n *Node) prettyPrint(prefix string) {
	pretabs := prefix + "\t"

	if n.Kind == URLNode {
		fmt.Printf(
			"%sNodeType %v->(url:%s), (alt:%s))\n",
			pretabs,
			n.Kind,
			n.URLData.URL,
			n.URLData.AltText,
		)
	} else if n.Kind == MetadataNode {
		fmt.Printf(
			"%sNodeType %v->[k:%s], [v:%s])\n",
			pretabs,
			n.Kind,
			n.Metadata.Key,
			n.Metadata.Value,
		)
	} else {
		fmt.Printf("%sNodeType %v->%s\n", pretabs, n.Kind, n.StringValue)
	}

	for _, child := range n.Children {
		child.prettyPrint(pretabs)
	}
}

func newNode(kind NodeType) *Node {
	return &Node{
		Kind:        kind,
		StringValue: "",
		Children:    []*Node{},
	}
}
