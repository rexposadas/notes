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

	// nodes in this channel are printed rightaway
	currentLevel := make(chan *Node, 1000)

	// notes in this channel will be printed after
	// a new line is generated.
	nextLevel := make(chan *Node, 1000)

	// Let's ready the current node for printing
	currentLevel <- n // root of the tree

	for len(currentLevel) > 0 {

		n := <-currentLevel
		fmt.Print(n.Value)

		if n.Left != nil {
			nextLevel <- n.Left
		}
		if n.Right != nil {
			nextLevel <- n.Right
		}

		if len(currentLevel) == 0 {
			fmt.Println("")
			swap(currentLevel, nextLevel)
		}
	}
}

// swap puts the next level of nodes to the "currentLevel"
// channel to be printed.
func swap(currentLevel, nextLevel chan *Node) {

	for len(nextLevel) > 0 {
		n := <-nextLevel
		currentLevel <- n
	}

	return
}
