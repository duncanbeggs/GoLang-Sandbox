//Reverse a linked list
package main

import "fmt"

type listNode struct {
	nextPtr *listNode
	nodeVal int
}

func addNode(np *listNode, intVal int) {
	newNode := listNode{nil, intVal}
	if np.nextPtr == nil {
		np.nextPtr = &newNode
		return
	} else {
		addNode(np.nextPtr, intVal) //recursive call
	}
	return //don't think this code will ever execute
}

func printList(np *listNode) {
	if np.nextPtr == nil { //if we get to end of the list
		fmt.Println(np.nodeVal)
		return
	} else {
		fmt.Printf("%d, ", np.nodeVal)
		printList(np.nextPtr) //recursive call, go to next node
	}
}

func findNode(np *listNode, intVal int) bool {
	if (np.nextPtr == nil) && (np.nodeVal != intVal) {
		//reached end of list without finding intVal
		return false
	} else if (np.nextPtr == nil) && (np.nodeVal == intVal) {
		//Got to end of list and found the node.
		return true
	} else if (np.nextPtr != nil) && (np.nodeVal != intVal) {
		//Did not find val but still not at end of list. Recurse onwards.
		return findNode(np.nextPtr, intVal)
	} else if (np.nextPtr != nil) && (np.nodeVal == intVal) {
		return true
		//Found val in middle of list.
	}
	//if we get to here something went wrong
	return false //insists on this being here
}

//non-recursive delete function
func deleteNode(firstNodePtr *listNode, intVal int) *listNode {
	var tempPtr *listNode
	var prevPtr *listNode
	prevPtr = nil                  //remove this later to test if this line is unnecessary
	currentNodePtr := firstNodePtr //storing this because np will change as it iterates
	for currentNodePtr != nil {
		if currentNodePtr.nodeVal == intVal { //if we find a match proceed to determine where in the list it is
			if prevPtr == nil { //case if we need to delete first node
				tempPtr = currentNodePtr.nextPtr
				currentNodePtr.nextPtr = nil
				return tempPtr //return pointer to next node in list
			} else if prevPtr != nil && currentNodePtr.nextPtr != nil { //case if we are in the middle of the list
				prevPtr.nextPtr = currentNodePtr.nextPtr
				currentNodePtr = nil //think this effectively deletes it
				return firstNodePtr  //head does not change
			} else if prevPtr != nil && currentNodePtr.nextPtr == nil { //case if we are at the end of the list
				prevPtr.nextPtr = nil
				return firstNodePtr
			}
		}
		prevPtr = currentNodePtr                //save current np
		currentNodePtr = currentNodePtr.nextPtr //iterate forward
	}
	return firstNodePtr //if iterated through whole loop then nothing was deleted so return false
}

func reverseList(firstNodePtr *listNode) *listNode {
	var prevPtr *listNode
	prevPtr = nil
	var tempPtr *listNode
	currentNodePtr := firstNodePtr
	for currentNodePtr != nil {
		if currentNodePtr == firstNodePtr { //case for first node. Could also test by using prevPtr == nil
			prevPtr = currentNodePtr
			tempPtr = currentNodePtr.nextPtr
			currentNodePtr.nextPtr = nil //first node has now become the last node
			currentNodePtr = tempPtr     //advance loop forward one node
		} else if prevPtr != nil && currentNodePtr.nextPtr != nil { //case for node being in middle of list
			tempPtr = currentNodePtr.nextPtr
			currentNodePtr.nextPtr = prevPtr //point to the left
			prevPtr = currentNodePtr
			currentNodePtr = tempPtr //advance loop forward one node
		} else if currentNodePtr.nextPtr == nil { //case we are at the end of the list
			currentNodePtr.nextPtr = prevPtr //point our last node to the left
			return currentNodePtr            //return as the new head of list
		}
	}
	//	fmt.Println("possible bug - reverseList loop exited without returning")
	return firstNodePtr //if we get to here something unexpected happened
}

func main() {
	n := 1
	firstNode := listNode{nil, 0}
	nPtr := &firstNode
	for i := 1; i < 50; i++ {
		n = (3 * i) + (i / 2)
		addNode(nPtr, n)
	}
	printList(nPtr)

	var userInput1, userInput2 int
	for userInput1 != -1 {
		fmt.Printf("Find(1), Delete(2), Reverse(3), Print(4), or Quit(-1)?> ")
		fmt.Scan(&userInput1)
		switch userInput1 {
		case 1:
			for userInput2 != -1 {
				fmt.Printf("Num to find?(-1 to quit)> ")
				fmt.Scan(&userInput2)
				if findNode(nPtr, userInput2) {
					fmt.Println("Found value ", userInput2)
				} else {
					fmt.Println("Did not find value ", userInput2)
				}
			}
			userInput2 = 0 //set back to 0 so that next use of this var doesn't immediately cause loop exit
		case 2:
			for userInput2 != -1 {
				fmt.Printf("Num to delete?(-1 to quit> ")
				fmt.Scan(&userInput2)
				nPtr = deleteNode(nPtr, userInput2)
			}
			userInput2 = 0
		case 3:
			nPtr = reverseList(nPtr)
			fmt.Printf("List reverseed!\n")
		case 4:
			printList(nPtr)
		}
	}
}
