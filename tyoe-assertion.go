package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)  // hello true

	f, ok := i.(int)
	fmt.Println(f, ok) //0 false

	f = i.(int)  // panic: interface conversion: interface {} is string, not int
	fmt.Println(f)
}
/*
类型断言：t := i.(T)
该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。
 */