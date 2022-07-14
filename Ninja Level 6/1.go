package main
			 
			 
import "fmt" 
			 
			 
func main() {
		
	a := foo();
	var b string
	var c int
	b,c = bar();

	fmt.Println(a,b,c);
			 
}

func foo() int {
	return 5;
}

func bar() (string , int ){
	return "deepak" , 22 ;
}