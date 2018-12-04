package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 101; i++ {
		sum += i
	}
	fmt.Printf("the sum is %d", sum)
}

/*
初始化语句：在第一次迭代前执行(可选）
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行(可选）
例如：
for sum < 100 {...}

条件表达式也省略时无限循环
例如：
for {...}
 */