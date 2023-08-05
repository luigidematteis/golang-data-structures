package main

import (
	"fmt"
	"learning/arrays"
	"learning/trees"
)

func main() {
	arrays.CallMethodsInPackageArray(false)

	n := &trees.Node{Value: 3, Left: nil, Right: nil}
	n.Left = &trees.Node{Value: 10, Left: nil, Right: nil}
	n.Right = &trees.Node{Value: 26, Left: nil, Right: nil}

	b := trees.Bst{Root: n, Len: 3}
	b.Add(9)
	b.Add(2)
	b.Add(8)
	b.Add(56)
	b.Add(58)
	fmt.Println(b)
}
