package main
			 
			 
import "fmt" 
		
func factorial(n int) int{
	if(n <= 1) {
		return n;
	}

	return n * factorial(n-1)
}
			 
func main() {
		
	// * Clousre

	{
		y := 3
		fmt.Println(y)
	}

	// fmt.Println(y) // y cannot be printed here

	/* a := incrementor();
	fmt.Println(a())
	fmt.Println(a())
	b := incrementor();
	fmt.Println(b());
	fmt.Println(a())
	fmt.Println(b()); */

	// * Recursion

	fmt.Println(factorial(5))


}

func incrementor() func() int{
	var x int
	return func() int{
		x++;
		return x;
	}
}