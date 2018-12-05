package main

import "fmt"

type Person struct {
	name string
}
//(p Person) is receiver, between func and methods name
func (p Person) info() string {
	return "the person's info is " + p.name
}

// pointer receiver to modify value
// if use value receiver, the output is not changed
func (p *Person) modify(s string) {
	p.name = s
}

func main() {
	p := Person{"lihua"}
	p.modify("hacker")
	fmt.Println(p.info())
}