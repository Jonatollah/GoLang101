package main

import(
	"fmt"
)

func main(){
	foo()
	a,b := mouse("ian","fleming")
	fmt.Println(a)
	fmt.Println(b)
}
// func (r receiver) identifier(parameters) (return(s)) { ...  }
func foo(){
	fmt.Println("hello  from foo")

}
func bar(s string){
	fmt.Println("Hello,",s)
}
func mouse(fn string, ln string)(string,bool){
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := true
	return a, b
}
