package main
			 
			 
import "fmt" 
		
type person struct{
	first string
	last string
	age int
}

type secretagent struct{
	person
	license_to_kill bool
}
			 
func main() {

	p1 := person{
		first: "James",
		last: "Bond",
		age: 34,
	}


	p2 := person{
		first: "James",
		last: "Bond the second",
	}

	fmt.Println(p1);
	fmt.Println(p2.first);

	// * Embedding Structs

	sa1 := secretagent{
		person: p1,
		license_to_kill: false,
	}

	fmt.Println(sa1);
	fmt.Println(sa1.first); // inner type gets promoted to the outer type


	// * Anonmymous Structs

	as := struct{
		class string
	}{
		class: "Does not exist in go",
	}

	fmt.Println(as);
			 
			 
}