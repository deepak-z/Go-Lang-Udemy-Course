package main

import (
	"encoding/json"
	"fmt" 
)
			 
type person struct{
	First string `json:"First"`
	Last string `json:"Last"`
	Age int `json:"Age"`
}
func main() {
	
	s := `[{"First":"Deepak","Last":"Kumar","Age":22},{"First":"D","Last":"K","Age":35}]`
	fmt.Println(s)
			
	bs := []byte(s)

	fmt.Printf("%T\n", s);
	fmt.Printf("%T\n", bs);

	//people := []person{}
	var people []person
	
	err := json.Unmarshal(bs,&people)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(people)
	}

	//fmt.Println(bs);
}