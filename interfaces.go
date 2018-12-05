// An interface is a collection of methods
package main

import "fmt"

// declare interface
type geometry interface {
	area() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g)
}

func main() {
	r := rectangle{12, 6}
	measure(r)
}
