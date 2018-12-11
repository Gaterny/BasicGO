package main

import "fmt"

func main() {
	foo := func() string{
		return "hello foo"
	}

	fmt.Println(foo()) // hello foo

	bar := func() {
		fmt.Println("hello bar")
	}

	bar() // hello bar

 	// 后面加括号可以执行一个函数
	func() {
		fmt.Println("hello func")
	}()  // hello func
}
