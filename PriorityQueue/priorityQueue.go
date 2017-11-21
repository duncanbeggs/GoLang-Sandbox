//implement a priority queue using the array implementation of a heap
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type heap struct {
	pq [41]int
	size int
}

//GoLang doesn't have constructors so one best practice is to use a New() function to return an initialized struct
func newHeap() *heap {
	var h heap
	h.size = 0
	for j := 0; j < 41; j++ {
		h.pq[j] = -1
	}
	return &h
}


func bubbleUp(h *heap, index int) {
	if index == 0 {
		return
	}
	if (index % 2 == 1) {//in case of odd index parent is (index -1)/2. Swap child and parent
		if (h.pq[index] < h.pq[(index-1)/2]) {//if child is less than parent swap them
			temp := h.pq[index]
			h.pq[index] = h.pq[(index-1)/2]
			h.pq[(index-1)/2] = temp
		}
		bubbleUp(h, ((index-1)/2) )//bubble up to next level
	} else if (index % 2 == 0) {//in case of even index (right child) the parent is at (index-2)/2. Swap child and parent
		if (h.pq[index] < h.pq[(index-2)/2]) {//if child is less than parent swap them
			temp := h.pq[index]
			h.pq[index] = h.pq[(index-2)/2]
			h.pq[(index-2)/2] = temp
		}
		bubbleUp(h, ((index-2)/2) )
	}
	return
}

func bubbleDown(h *heap, index int) {
	if (index > 19) {//if i is greater than 19 than we are at a leaf. Proceeding further down this code would cause out of bounds error
		return
	}
	if ( (h.pq[index] > h.pq[2*index+1]) && (h.pq[index] > h.pq[2*index+2]) ) {//bigger than both children
		if (h.pq[2*index+1] < h.pq[2*index+2]) {//if right child is bigger than left
			//swap right child for parent
			temp := h.pq[2*index+2]
			h.pq[2*index+2] = h.pq[index]
			h.pq[index] = temp
			bubbleDown(h, (2*index+2))
		} else {//else left child is equal to or greater than parent
			temp := h.pq[2*index+1]
			h.pq[2*index+1] = h.pq[index]
			h.pq[index] = temp
			bubbleDown(h, (2*index+1))
		}
	} else if (h.pq[index] > h.pq[2*index+2]) {//bigger than just right child
			temp := h.pq[2*index+2]
			h.pq[2*index+2] = h.pq[index]
			h.pq[index] = temp
			bubbleDown(h, (2*index+2))
	} else if (h.pq[index] > h.pq[2*index+1]) {//bigger than just left child
			temp := h.pq[2*index+1]
			h.pq[2*index+1] = h.pq[index]
			h.pq[index] = temp
			bubbleDown(h, (2*index+2))
	}
	return
}


func poll(h *heap) int {
	n := h.pq[0]
	//make h.pq[0 and bubble down
	h.pq[0] = h.pq[h.size-1]
	h.pq[h.size-1] = -1
	h.size--
	bubbleDown(h, 0)
	return n
}

//this is slightly inefficient, fix loop
func insert(n int, h *heap) bool {
	if (*h).size < 41 {
		for i := 0; i < 41; i++ {
			if ((*h).pq[i] == -1) {
				(*h).pq[i] = n
				(*h).size++
				bubbleUp(h, (h.size-1))//h.size-1 is the value of the index we want to start bubbling up from
				return true
			}
		}
	}
	return false
}

//take an int to delete
func deleteElement(h *heap, n int) bool {
	for i := 0; i < h.size; i++ {

		if (n == h.pq[i]) {//check if match
			//first swap value to be deleted with the last value (h.size-1)
			temp := h.pq[h.size-1]//in this case mostly temp
			h.pq[i] = temp
			h.pq[h.size-1] = -1//put negative one in the last spot to mark it as empty
			h.size--//decrement size by 1
			if ( i == 0 ) {
				bubbleDown(h, 0)
				return true
			}
			if ( i % 2 == 0) {//check for even or odd to determine how to calc parent
				if (temp < h.pq[(i-2)/2]){//if swapped element is smaller than parent bubble up
					bubbleUp(h, i)
					return true
				}
				if (i < 20) {//If i is greater than 20 then we would be at the bottom of tree so bubbling down not needed
					if (temp > h.pq[2*i+1] || temp > h.pq[2*i+2]) {//if swapped element is larger than parent bubbld down
						bubbleDown(h, i)
						return true
					}
				}
			} else { //odd index, different parent calculation
				if (temp < h.pq[(i-1)/2]){//if swapped element is smaller than parent bubble up
					bubbleUp(h, i)
					return true
				}
				if ( i < 20) {
					if (temp > h.pq[2*i+1] || temp > h.pq[2*i+2]) {//if swapped element is larger than parent bubble down
						bubbleDown(h, i)
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	myHeap := newHeap()
	fmt.Println(myHeap)
	var n int
	for i := 0; i < 41; i++ {
		n = (r.Int() % 128)
		insert(n,myHeap)
		time.Sleep(time.Millisecond * 40)
		fmt.Printf("Inserting: %d||", n)
		fmt.Println(myHeap)
	}
	var x int
	fmt.Printf("Delete what?>")
	fmt.Scan(&x)
	deleteElement(myHeap, x)
	fmt.Println("\n", myHeap)
	fmt.Printf("Delete what?>")
	fmt.Scan(&x)
	deleteElement(myHeap, x)
	fmt.Println("\n", myHeap)
	fmt.Printf("Delete what?>")
	fmt.Scan(&x)
	deleteElement(myHeap, x)
	fmt.Println("\n", myHeap)
	fmt.Printf("Delete what?>")
	fmt.Scan(&x)
	deleteElement(myHeap, x)
	fmt.Println("\n", myHeap)
	fmt.Printf("Delete what?>")
	fmt.Scan(&x)
	deleteElement(myHeap, x)
	fmt.Println("\n", myHeap)
	fmt.Printf("Delete what?>")
	fmt.Scan(&x)
	deleteElement(myHeap, x)
	fmt.Println("\n", myHeap)
	fmt.Printf("Delete what?>")
	fmt.Scan(&x)
	deleteElement(myHeap, x)
	fmt.Println("\n", myHeap)
}
