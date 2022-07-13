package main
			 
			 
import "fmt" 
			 
			 
func main() {
			
	a := [5]int{1,2,3,4,5};

	for i := range a{
		fmt.Println(i,a[i]);
	}

	fmt.Printf("%T\n",a);


			 
}