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
	isEmpty bool
	//might put some more stuff in here later
}

func nodeInsert(intVal int, t *Tree) bool {
	if (*t).isEmpty == true {
		newNode := treeNode{nil, nil, intVal}
		t.root = &newNode
		(*t).isEmpty = false
	} else {
		recNodeInsert(intVal, (*t).root)
	}
	return true
}

func recNodeInsert(intVal int, tn *treeNode) {
	if tn == nil {
		fmt.Printf("|tn==nil Node Inserted|\n")
		newNode := treeNode{nil,nil,intVal}
		tn = &newNode
		return
	}

	if tn.nodeVal < intVal {//NOTE: switching from (*tn).nodeVal syntax. These are equivalent in Go
		fmt.Printf("inserting %d right of %d",intVal,tn.nodeVal)
		recNodeInsert(intVal, tn.rightPtr)
	} else if tn.nodeVal > intVal {
		fmt.Printf("inserting %d left of %d", intVal, tn.nodeVal)
		recNodeInsert(intVal, tn.leftPtr)
	} else {//in case of tn.nodeVal == intVal then do nothing
		fmt.Printf("|nodeVal == intVal|")
		return
	}
}

func printInOrder (tn *treeNode) {
	if tn == nil {
		fmt.Println("Tree is empty")
		return
	}

	if (tn.leftPtr == nil && tn.rightPtr == nil) {
		fmt.Printf("%d, ", tn.nodeVal)
	} else if tn.leftPtr == nil && tn.rightPtr != nil {
		printInOrder(tn.rightPtr)
		fmt.Printf("%d, ", tn.nodeVal)
	} else if tn.leftPtr != nil && tn.rightPtr == nil {
		fmt.Printf("%d, ", tn.nodeVal)
		printInOrder(tn.leftPtr)
		return
	} else if tn.leftPtr != nil && tn.rightPtr != nil {
		printInOrder(tn.leftPtr)
		fmt.Printf("%d, ", tn.nodeVal)
		printInOrder(tn.rightPtr)
		return
	}
	return
}


func main() {
	printWelcome()
	fmt.Println("Tree Ready! What would you like to do?")
	fmt.Printf("1 - Insert int\n2 - Delete int\n3 - Check for int\n4 - Print Tree\n5 - Quit\n> ")
	x := 0
	n := 0
	myTree := Tree{nil, true}
	treePtr := &myTree
	for  x != 5 {
		fmt.Scan(&x)
		switch x {
		case 1:
			fmt.Printf("Insert int> ")
			fmt.Scan(&n)
			nodeInsert(n, treePtr)
		case 2:
			fmt.Printf("Delete int> ")
			fmt.Scan(&n)
		case 3:
			fmt.Printf("Search for int> ")
			fmt.Scan(&n)
		case 4:
			fmt.Printf("Print in-order(1), pre-order(2) or post-order(3)?>")
			fmt.Scan(&n)
			switch n {
				case 1:
					fmt.Printf("In-Order: ")
					printInOrder((*treePtr).root)
				case 2:
					fmt.Printf("pre-order")
				case 3:
					fmt.Printf("post-order")
				default:
					fmt.Println("Option not valid")
			}
		case 5:
			fmt.Println("ByeBye!!")
		default:
			fmt.Printf("You fucked up. INTEGER 1-5!> ")
			fmt.Scan(&x)
		}
		if x != 5 {
			fmt.Printf("Anything else?\n> ")
		}
	}
}

