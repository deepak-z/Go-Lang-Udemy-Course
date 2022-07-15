package main

import (
	"encoding/json"
	"fmt" 
	"os"
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

	 err := json.NewEncoder(os.Stdout).Encode(users)
	 if err !=  nil {
		fmt.Println("We did something wrong")
	 }
			 
}