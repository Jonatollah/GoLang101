package main
import (
	"fmt"
)
func main(){
	for i=0;i<=100; i++ {
		fmt.Println(i)
	}
	for j :=0; j<=10; j++{
		for d :=0; d<3; d++{
			fmt.Println(j)
		}
	}
	for x<10{
		fmt.Println(x)
		x++
	}
	m:=1
	for {
		if m>9{
			break
		}
		fmt.Println(m)
		m++
	}
	fmt.Println("done.")
	for n :=33; n<= 122; n++{
		fmt.Printf("%v\t%#x\t%#U\n",n,n,n)
	}
}