import ( 
	"fmt"
)
type person struct {
	first string
	last string
}
type secretAgent struct {
	person
	ltk bool
}
func (s secretAgent) speak(){
	fmt.Println("I am", s.first, s.last)
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}
// order: keyword, identifier, type
//interface  says hey baby, if you got this method
//then you're my type
type human interface {
	speak()
}
func bar(h human){
	switch h.(type){
	case person:
		fmt.println("I was passed into bar as person", h.(person).first)
	case secretAgent:
		fmt.println("I was passed into bar as secretAsgent", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar",h)
}

func main(){
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}
	p1:= person{
		first: "Dr.",
		last: "yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}