package hello

import "fmt"

func init() {
	fmt.Println("using init function")
}

func foo() string {
	return "using foo function"
}
