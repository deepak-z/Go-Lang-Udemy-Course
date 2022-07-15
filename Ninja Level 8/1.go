package main

import (
	"encoding/json"
	"fmt" 
)
	
type user struct{
	LoginId string
	Password string
}
			 
func main() {

	 u1 := user{
		LoginId: "DEEP",
		Password: "I dont care",
	 }


	 u2 := user{
		LoginId: "KUMA",
		Password: "I still don't care",
	 }

	 users := []user{u1,u2};

	 fmt.Println(users)

	 bs,err := json.Marshal(users)

	 if (err != nil){
		fmt.Println(err)
	 }else{
		fmt.Println(string(bs))
	 }

			 
}