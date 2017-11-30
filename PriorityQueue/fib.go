package main

import "fmt"

func fibNum(x int) {
	var a uint64
	var b uint64
	a = 0
	b = 1
	fmt.Println(0)
	for i := 0; i < x; i++ {
		if ( i % 2 == 0 ) {
			a = a + b
			fmt.Println(a)
		} else {
			b = a + b
			fmt.Println(b)
		}
	}
}




func main() {
	fmt.Printf("Welcome to the cool fibonacci sequence\n")
	fmt.Printf("Calculate how many fibonacci numbers?> ")
	var x int
	fmt.Scan(&x)
	fibNum(x)
}

