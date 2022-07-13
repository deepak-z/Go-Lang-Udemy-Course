package main
			 
			 
import "fmt" 
			 
			 
func main() {
	
	x := make([]string,50,50);

	fmt.Println(x);

	y := []string{"a","b","c","d"};

	for i := range y {
		x[i] = y[i];
	}
	fmt.Println(y);
	fmt.Println(x);
			 
}