package main
import("fmt")
var z = 22
var y int =0
func main(){
	x :=42
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 +24 +y
	foo()
	fmt.Println(y)
	d:="||||||"
	fmt.Println(d)
	fmt.Printf("%T\n",y)

	fmt.Printf("%b\n",y)
	fmt.Printf("%x\n",y)
	y=911
	fmt.Printf("%#x\n",y)
	s:= fmt.Sprintf("%#x\t%b\t%x\n",y,y,y)
	fmt.Println(s)
}

func foo(){
fmt.Println("again",z)

}