package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
/*
Go 从上至下依次判断，只运行选定的 case，而非之后所有的 case。
Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
除非以 fallthrough 语句结束，否则分支会自动终止。
 */
