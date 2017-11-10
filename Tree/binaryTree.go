//Writing a custom implementation of a binary tree
//All in one file for now. Will break out to multiple files after figuring out how to do so.
package main

import (
	"fmt"
)

type treeNode struct {
	leftPtr *treeNode
	rightPtr *treeNode
	nodeVal int
}

type Tree struct {
	root *treeNode
	//might put some more stuff in here later
}

func newTree() *Tree {
	var t *Tree
	return t
}

func main() {
	printWelcome()
	fmt.Println("Tree Ready! What would you like to do?")
	fmt.Printf("1 - Insert int\n2 - Delete int\n3 - Check for int\n4 - Print Tree")


	fmt.Println("One day I will be a big beautiful tree!")
}
