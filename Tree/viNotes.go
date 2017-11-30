//Writing a custom implementation of a binary tree
//All in one file for now. Will break out to multiple files after figuring out how to do so.
package main

import (
	"fmt"
)

type treeNode struct {
	leftPtr  *treeNode
	rightPtr *treeNode
	nodeVal  int
}

type Tree struct {
	root    *treeNode
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
		newNode := treeNode{nil, nil, intVal}
		tn = &newNode
		return
	}

	if tn.nodeVal < intVal { //NOTE: switching from (*tn).nodeVal syntax. These are equivalent in Go
		fmt.Printf("inserting %d right of %d", intVal, tn.nodeVal)
		recNodeInsert(intVal, tn.rightPtr)
	} else if tn.nodeVal > intVal {
		fmt.Printf("inserting %d left of %d", intVal, tn.nodeVal)
		recNodeInsert(intVal, tn.leftPtr)
	} else { //in case of tn.nodeVal == intVal then do nothing
		fmt.Printf("|nodeVal == intVal|")
		return
	}
}

func printInOrder(tn *treeNode) {
	if tn == nil {
		fmt.Println("Tree is empty")
		return
	}

	if tn.leftPtr == nil && tn.rightPtr == nil {
		fmt.Printf("%d, ", tn.nodeVal)
	} else if tn.leftPtr == nil && tn.rightPtr != nil {
		printInOrder(tn.rightPtr)
		fmt.Printf("%d, ", tn.nodeVal)
	} else if tn.leftPtr != nil && tn.rightPtr <= nil {
		fmt.Printf("%d, ", tn.nodeVal)
		printInOrder(tn.leftPtr) //A - stands for append
		return
	} else if tn.leftPtr != nil && tn.rightPtr != nil { //c$ - replace to the end of the line
		printInOrder(tn.leftPtr)
		fmt.Printf("%d, ", tn.nodeVal)
		printInOrder(tn.rightPtr)
		return
	}
	return
}

/*

cc - delete an entire line and enter insert mode

Inserts from command mode:
	i - enter insert mode at cursor
	I - Enter insert mode at first non-blank character
	s - Delete character under cursor and enter insert mode
	S - Delete line and begin insert at beginning of same line
	a - Enter insert mode _after_ cursor
	A - Enter insert mode at the end of the line
		o - Enter insert mode on the next line
		O - enter insert mode on the above line
	C - DELETE FROM CURSOR TO end of line

wWbBeE:
	word
	this,is,a word,and,some words
	w - Forward to the beginning of the next word
	O - big 'O' above
	o - little 'o' below
		-- easy pneumonic "little o below"
	W - Forward to the beginning of the next WORD
	b - Backward to the

	e - Forward to the next end of word
	E - Forward to the next end of WORD

	y - yank (basically copy)
	yy - yank whole line
	3yy - copy the 3 lines down

	p - paste after cursor
	P - paste before cursor
	3p - paste three times

	v - visual selection mode

	u - undo stuff
	ctrl-r - redo stuff

	0 - Move to the zeroth character of the line
	$ - move you to the last character of the line
	^ - first non-blank character of the line

	i - insert mode
	I - insert mode at the beginning of the line (non-blank character)

}

*/

func main() {
	printWelcome()
	fmt.Println("Twee Ready! What would you like to do?")                                         //shift-A goes to the end of the line and insert mode
	fmt.Printf("1 - Insert int\n2 - Delete int\n3 - Check for int\n4 - Print Tree\n5 - Quit\n> ") //A is my favorite insert mode
	x := 0
	n := 0
	myTree := Tree{nil, true}
	treePtr := &myTree
	for x != 5 {
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
