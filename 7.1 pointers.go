package main
			 
			 
import "fmt" 
			 
			 
func main() {
		
	a := 42
	fmt.Println("a value :" , a)
	fmt.Println("a address :" , &a) // & gives you the address 

	fmt.Printf("Type of a %T\n",a)
	fmt.Printf("Type of &a %T\n",&a)

	var b *int
	b = &a
	fmt.Println(b)
	fmt.Printf("Type of  b %T\n",b)
	fmt.Println("Value of  *b ",*b) // gives you the value stored at an address when you have the address
	fmt.Println("Works also like this   *&a",*&a)

	*b = 43
	fmt.Println(a) // The value of a has changed 
			 
}