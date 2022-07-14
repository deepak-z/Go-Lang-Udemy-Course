package main
			 
			 
import "fmt" 
			 
			 
func main() {
		
	f := foo();
	w := f(45);
	fmt.Println(w)
			 
}


func foo() func(x int) int{
	return func(x int) int{
		return x+10;
	}
}