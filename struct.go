package main

import "fmt"

type Student struct {
	name string
	age int
	email string
	phone int
}

func main() {
	s:= Student{"lihua", 18, "lihua@yahoo.com", 5421753}
	s.age = 20
	fmt.Println(s.age) // 20
}