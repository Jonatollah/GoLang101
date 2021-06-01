package main

import ("fmt")
type person struct{
	first string
	last string
	age int
}
type secretAgent struct {
	person
	ltk bool
}
func main(){
	p1:=person{
		first:"james",
		last:"Bond",
		age:32,
	}
	p2 := secretAgent{
		person: person{
			first:"james",
			last:"bond",
			age:32,
		},
		ltk:true,
	}
	//anonymous struct
	p3:=struct{
		first string
		last string
		age int
	}{
		first:"James",
		last:"Bond",
		age: 32,
	}
	fmt.Println(p1)
	fmt.Println(p1.first,p1.last)
	fmt.Println(p2)
	fmt.Println(p2.first,p2.last,p2.age,p2.ltk)
}