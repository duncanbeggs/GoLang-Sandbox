//Writing a custom implementation of a binary tree
//All in one file for now. Will break out to multiple files after figuring out how to do so.
package main

import (
	"fmt"
	"time"
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
	fmt.Println("***-Welcome to Duncan's Binary Tree!!-***")
	fmt.Println("     *    *    ")
	fmt.Println("*  * *   *   * ")
	fmt.Println("**  **  ** **  ")
	fmt.Println("  * *    **    ")
	fmt.Println("  **    *   *  ")
	fmt.Println("  **   **   ** ")
	fmt.Println("    **  *  **  ")
	fmt.Println(" **  ** ***    ")
	fmt.Println("   **** **     ")
	fmt.Println("     ****      ")
	fmt.Println("     ****      ")
	fmt.Println("     ****      ")
	fmt.Print("New Tree Instatiating")
	time.Sleep(1200 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(120 * time.Millisecond)
	fmt.Print("..")
	time.Sleep(480 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(200 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(200 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(160 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(80 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(600 * time.Millisecond)
	fmt.Print("..")
	time.Sleep(160 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(50 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(90 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(800 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(1800 * time.Millisecond)
	fmt.Print("..")
	time.Sleep(50 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(50 * time.Millisecond)
	fmt.Print("...")
	time.Sleep(50 * time.Millisecond)
	fmt.Print(".....\n")
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Tree Ready! What would you like to do?")



	fmt.Println("One day I will be a big beautiful tree!")
}
