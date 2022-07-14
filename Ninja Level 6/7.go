package main
			 
			 
import "fmt" 
	
var g func(x int)
			 
func main() {
			
	f := func (x int)  {
		fmt.Println(x)
	}

	g = f

	f(42)
	g(43)
}