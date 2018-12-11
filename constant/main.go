package main

import "fmt"

// 常量一般是不变的，不会修改的
const (
	one  = iota + 1
	two
	three
)

func main() {
	fmt.Println(one, two, three)
}
