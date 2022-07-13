package main
			 
			 
import "fmt" 
			 
			 
func main() {
		
	x := []int{1,2,3,4,5};
	fmt.Println(x);
	x = append(x, 6);
	fmt.Println(x);
	x = append(x, 8,9,10);
	fmt.Println(x);
	a := []int{11,12,13};
	x = append(x, a...);
	fmt.Println(x);
			 
}