package main
			 
			 
import "fmt" 
			 
			 
func main() {
	
	// * Anonymous Self Executing Functions 
	/* foo()

	func(){
		fmt.Println("Anonymous Funtion Ran")
	}()

	func(x int){
		fmt.Println("The meaning of Life:",x)
	}(42) */

	// * Func Expression

	/* f := func(x int){
		fmt.Println("This is my first func expression",x)
	}

	f(42) */

	// * Returning Function

	/* s1 := returing_func()
	fmt.Println(s1)

	srf := self_returning_func()
	fmt.Printf("%T\n",srf);
	w := srf();
	fmt.Println(w)

	// ? can also be done like this

	fmt.Println(self_returning_func()()) */

	// * Call back

	xi := []int{1,2,3}
	s := sum(xi...)
	fmt.Println(s)

	ei := []int{1,2,3,100}
	ei_sum := even(sum,ei...)
	fmt.Println(ei_sum)



			 
}

func sum(xi ...int)int {
	total := 0
	for _,v := range xi{
		total += v
	}
	return total
}

// * passing a function as a argument

func even(f func(xi ...int) int , vi ...int) int {
	var yi []int
	for _,v := range vi{
		if(v % 2 == 0) {
			yi = append(yi, v)
		}
	}

	return f(yi ...)
}

func returing_func() string{
	s := "I don't care"
	return (s)
}

func self_returning_func() func() int{
	return func() int{
		return 445;
	}
}
func foo(){
	fmt.Println("foo ran")
}