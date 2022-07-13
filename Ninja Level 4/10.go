package main
			 
			 
import "fmt" 
			 
			 
func main() {
			 
	mp :=  map[string][]string{
		"Deepak Kumar" : {"I","do","not","care"},
		"Deepak Kumar from Futre" : {"I","still","do","not","care"},
	}	
	
	mp["woo hoo"] = []string{"This","is","just","random","stuff"}

	if i,ok := mp["woo hoo"]; ok{
		delete(mp,"woo hoo")
		fmt.Println(i)
	}

	for k,val := range mp{
		fmt.Println(k);
		for i,v := range val{
			fmt.Println(i,v)
		}
	}
}