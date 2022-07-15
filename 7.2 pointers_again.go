package main
			 
			 
import "fmt" 
			 
			 
func main() {
			 
		/* x := 20
		fmt.Println(x)
		foo(x)
		fmt.Println(x)	 */ 


		x := 20
		fmt.Println(x)
		foo(&x)
		fmt.Println(x)	
}

/* func foo(x int){
	fmt.Println(x)
	x = 43
	fmt.Println(x)
} */

func foo(x *int){
	fmt.Println(x)
	*x = 43
	fmt.Println(x)
}