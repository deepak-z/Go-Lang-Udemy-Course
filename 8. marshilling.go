package main

import (
	"encoding/json"
	"fmt" 
)

type person struct{
	First string   // * first letter of identifier should be capital for marshalling ot work
	Last string
	Age int
}
			 
			 
func main() {

	// * MARSHLING

	p1 := person{
		First: "Deepak",
		Last: "Kumar",
		Age: 22,
	}

	p2 := person{
		First: "D",
		Last: "K",
		Age: 35,
	}
	
	people := []person{
		p1,
		p2,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(people)

	bs ,err := json.Marshal(people)

	if(err != nil){
		fmt.Println(err)
	}else { 
		fmt.Println(string(bs))
	}
	

	// * UNMARSHLING
}