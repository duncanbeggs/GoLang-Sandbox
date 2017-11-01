package main

import "fmt"

func fibo(n int) (sum int) {
	if (n <= 2) {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	fmt.Println("Calculate fibonacci sequence to what term?")
	var n, x int
	fmt.Scan(&n)
	x = fibo(n)
	fmt.Printf("The %dth term is %d\n", n, x)
}

