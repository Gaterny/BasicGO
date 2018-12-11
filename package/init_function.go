package main

/*
init方法是始终会执行的
当我们只使用init方法，不想使用其他方法时可以采用。
 */
import (
	"fmt"
	_ "BasicGO/package/test"
)

func main() {
	fmt.Println("Do not use imported function")
}
