package main

import (
	"fmt"
)
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
	{
		y:=42
		fmt.Println(y)
	}
	a:= incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}
func foo(){fmt.Println("hello")
x++}
func incrementor() func() int{
	var x intreturn func() int {
		x++
		return x
	}
}