package main

import "fmt"

/*
the function return a function that return int
 */
func fib() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	f := fib()
	fmt.Println(f())  // 1
	fmt.Println(f())  // 2
	fmt.Println(f())  // 3
}
