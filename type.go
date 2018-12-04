/*
type 类型名字 底层类型
 */
package main

import "fmt"

//华氏温度和摄氏温度
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit{
	// Fahrenheit is type convert, not function call,
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius{
	return Celsius((f-32) *5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Println(boilingF)
	absoluteZeroF := FToC(CToF(AbsoluteZeroC))
	fmt.Println(absoluteZeroF)
	// differnt type can't be compared
	fmt.Println(boilingF == absoluteZeroF) //mismatched type
}

