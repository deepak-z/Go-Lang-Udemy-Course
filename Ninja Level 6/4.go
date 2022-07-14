package main
			 
			 
import "fmt" 

type person struct{
	first string
	last string
	age int
}

func (p person) speak(){
	fmt.Println("hey I am", p.first, "and I am ",p.age," years old");
}
			 
			 
func main() {
			 
	p := person{
		first: "Deepak",
		last: "Kumar",
		age: 22,
	}

	fmt.Println(p)
	p.speak()
}