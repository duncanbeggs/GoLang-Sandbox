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
	fmt.Printf("1 - Insert int\n2 - Delete int\n3 - Check for int\n4 - Print Tree\n5 - Quit\n> ")
	x := 0
	for  x != 5{
		fmt.Scan(&x)
		switch x {
		case 1:
			fmt.Printf("Insert int> ")
		case 2:
			fmt.Printf("Delete int> ")
		case 3:
			fmt.Printf("Search for int> ")
		case 4:
			fmt.Printf("Print in-order(1), pre-order(2) or post-order(3)?>")
		case 5:
			fmt.Println("ByeBye!!")
		default:
			fmt.Printf("You fucked up. INTEGER 1-4!> ")
			fmt.Scan(&x)
		}
	}
	fmt.Println("One day I will be a big beautiful tree!")
}
