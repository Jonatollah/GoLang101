package main
import("fmt")

const(
	x=42
y int =43
)
func main(){
	a :=(42==42)
	b :=(42<=43)
	c :=(42>=43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)
	l:= `here  is something
	as
	a
	raw string
	literal
	"you see"
	another thing`
	fmt.Prinln(l)
	fmt.Println(a,b,c,d,e,f)
}