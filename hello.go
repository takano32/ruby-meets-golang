package main

import (
	"C"
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

//export Init_hello
func Init_hello() {
	fmt.Println("hello.")
}

func main() {
}
