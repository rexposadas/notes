package main

// This sample code shows how to print a binary tree.
// Each row of level of the tree should be printed on a
// separate line.

import (
	"fmt"
)

func main() {

	// top node, considered  the first level
	root := &Node{
		Value: 1,
	}

	// level 2
	n2 := &Node{Value: 2}
	n3 := &Node{Value: 3}
	root.Left = n2
	root.Right = n3

	// level 3
	n4 := &Node{Value: 4}
	n2.Left = n4

	n5 := &Node{Value: 5}
	n3.Right = n5

	fmt.Println("Printing root node:")
	root.Print()

	fmt.Println("Printing node n3:")
	n3.Print()
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Print prints the entire tree from this node.
func (n *Node) Print() {

	if n == nil {
		return
	}

	q1 := make(chan *Node, 1000)
	q2 := make(chan *Node, 1000)

	q1 <- n // root of the tree

	for len(q1) > 0 {

		nn := <-q1
		fmt.Print(nn.Value)

		if nn.Left != nil {
			q2 <- nn.Left
		}
		if nn.Right != nil {
			q2 <- nn.Right
		}

		if len(q1) == 0 {
			fmt.Println("")
			swap(q1, q2)
		}
	}
}

func swap(q1, q2 chan *Node) {

	for len(q2) > 0 {
		n := <-q2
		q1 <- n
	}

	q2 = make(chan *Node)
	return
}
