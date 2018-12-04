/*
数组是一个由固定长度的特定类型元素组成的序列.
一个数组可以由零个或多个元素组成.
*/
package main

import "fmt"

//declare an array
var a[3]int

func main()  {
	b := [5]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 5}
	// 指定索引的值，其余为0
	d := [...]int{2: 4}

	fmt.Println(a)
	fmt.Println(d[2])
	fmt.Println(b[len(a)-1])

	//indices and elements
	for i, v := range c {
		fmt.Printf("%d %d\n", i, v)
	}
}