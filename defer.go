/*
defer调用函数，推迟调用的函数会立即求值并被压入一个栈中
外层函数返回前不会调用该函数，调用该函数时按照后进先出的顺序
 */
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
