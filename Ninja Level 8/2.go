package main

import (
	"encoding/json"
	"fmt" 
)

type user struct {
    LoginId string `json:"LoginId"`
    Password string `json:"Password"`
}
			 
func main() {
	
	s := `[{"LoginId":"DEEP","Password":"I dont care"},{"LoginId":"KUMA","Password":"I still don't care"}]`

	bs := []byte(s)

	var users []user

	err := json.Unmarshal(bs,&users)

	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println(users)
	}
			 
}