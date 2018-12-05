/*
Unlike functions, methods require receiver
 */
package main

import "fmt"

type Person struct {
	name string
}
//(p Person) is receiver, between func and methods name
func (p Person) info() string {
	return "the person's info is " + p.name
}

func main() {
	p := Person{"lihua"}
	fmt.Println(p.info())
}
