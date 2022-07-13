package main
			 
			 
import "fmt" 
			 
			 
func main() {
		
	x := []int{1,2,3,4,5};
	x = append(x[:2],x[3:]...);
	fmt.Println(x);
			 
}