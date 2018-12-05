/*
Slice（切片）代表变长的序列，序列中每个元素都有相同的类型。
一个slice由三个部分构成：指针、长度和容量。
内置的len和cap函数分别返回slice的长度和容量。
 */

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	//type, length, capacity
	b := make([]int, 5, 8)

	// add item into slice
	a = append(a, 5)
	b = append(b, 2, 3, 4)
	fmt.Println(a)  	// [1 2 3 4 5]
	fmt.Println((b))   	//[0 0 0 0 0 2 3 4]
}
