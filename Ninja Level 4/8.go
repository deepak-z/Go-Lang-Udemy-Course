package main
			 
			 
import "fmt" 
			 
			 
func main() {
			 
	mp :=  map[string][]string{
		"Deepak Kumar" : {"I","do","not","care"},
		"Deepak Kumar from Futre" : {"I","still","do","not","care"},
	}	 

	for k,val := range mp{
		fmt.Println(k);
		for i,v := range val{
			fmt.Println(i,v)
		}
	}
}