package main

import "fmt"

func switchType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case bool:
		fmt.Println("type is bool")
	}
}

func main() {
	switchType("hello")
	switchType(true)
}