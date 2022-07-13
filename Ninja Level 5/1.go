package main
			 
			 
import "fmt" 

type person struct{
	first string
	last string
	ficf []string
}
			 
			 
func main() {
		
	p1 := person{
		first: "Deepak",
		last: "Kumar",
		ficf: []string{"pista","butterscotch"},
	}

	fmt.Println(p1)

	for i,v := range p1.ficf{
		fmt.Println(i,v)
	}
			 
}