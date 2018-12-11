package main

import "fmt"

// return func() int
func foo() func() int {
	return func() int {
		return 100
	}
}

func main()  {
	bar := foo()
	fmt.Printf("%T", bar)  // func() int
	fmt.Println(bar()) // 100
}
