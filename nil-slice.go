package main

import "fmt"

func main() {
	var s[]int
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
}
/*
nil切片长度和容量皆为0，且没有底层数组，不能直接赋值
 */