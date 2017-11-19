//implement a priority queue using the array implementation of a heap
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type heap struct {
	pq [40]int
	size int
}

//GoLang doesn't have constructors so one best practice is to use a New() function to return an initialized struct
func newHeap() *heap {
	var h heap
	h.size = 0
	for j := 0; j < 40; j++ {
		h.pq[j] = -1
	}
	return &h
}

func bubbleUp(h *heap, index int) {
	//so size-1 is where our highest value is stored
	//that makes sense. Think about inserting your first value 7 into pq[0], the size will be 1
	if index == 1 {
		return
	}
	//if child is smaller than its parent
	if ((index-1)
	if h.pq[index-1] <= h.pq[index-2] {
		//swap h.pq[index-1] and h.pq[index-2]
		temp := h.pq[index-1]
		h.pq[index-1] = h.pq[index-2]
		h.pq[index-2] = temp
		bubbleUp(h, (index-1))
	}
	return
}

func bubbleUp2(h *heap, index int) {
	if index == 0 {
		return
	}
	if (index % 2 == 1) {//in case of odd index parent is (index -1)/2. Swap child and parent
		temp := h.pq[index]
		h.pq[(index-1)/2]

func bubbleDown(h *heap, index int) {
	if (index == h.size-1) {
		return
	}
	if h.pq[index] > h.pq[index+1] {
		temp := h.pq[index]
		h.pq[index] = h.pq[index+1]
		h.pq[index+1] = temp
		bubbleDown(h, index + 1)
	}
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

func insert(n int, h *heap) bool {
	if (*h).size < 40 {
		for i := 0; i < 40; i++ {
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

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	myHeap := newHeap()
	fmt.Println(myHeap)
	var n int
	for i := 0; i < 40; i++ {
		n = (r.Int() % 128)
		insert(n,myHeap)
	}
	fmt.Println("")
	fmt.Println(myHeap, "\n")
	fmt.Println(poll(myHeap))
	fmt.Println(myHeap)
	fmt.Println(myHeap, "\n")
	fmt.Println(poll(myHeap))
	fmt.Println(myHeap)
	fmt.Println(myHeap, "\n")
	fmt.Println(poll(myHeap))
	fmt.Println(myHeap)
}

