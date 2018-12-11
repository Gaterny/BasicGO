package main

import "fmt"

// 交换两个变量的值
func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	a := 10
	b := 20
	a, b = swap(a, b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
