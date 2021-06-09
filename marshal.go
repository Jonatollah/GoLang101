package main

import (
	"fmt"
	"encoding/json"
)
type person struct {
	first string
	last string
	age int
}
func main(){
	p1: =person{
		first:"James",
		last:"Bond",
		age:32,
	}
	p2 := person{
		first: "Miss",
		last: "Moneypenny".
		age: 27,
	}
	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
}