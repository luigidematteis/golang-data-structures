package trees

import (
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Bst struct {
	Root *Node
	Len  int
}

func (n Node) String() string {
	return strconv.Itoa(n.Value)
}

func (b Bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b *Bst) Add(value int) {
	b.AddToNode(b.Root, value)
	b.Len++
}

func (b *Bst) AddToNode(root *Node, value int) *Node {
	if root == nil {
		return &Node{
			Value: value,
		}
	}

	if value < root.Value {
		root.Left = b.AddToNode(root.Left, value)
	} else {
		root.Right = b.AddToNode(root.Right, value)
	}

	return root
}

func (b Bst) inOrderTraversal(sb *strings.Builder) {
	b.nodeTraversal(sb, b.Root)
}

func (b Bst) nodeTraversal(sb *strings.Builder, root *Node) {
	if root == nil {
		return
	}
	b.nodeTraversal(sb, root.Left)
	sb.WriteString(root.String() + " ")
	b.nodeTraversal(sb, root.Right)
}
