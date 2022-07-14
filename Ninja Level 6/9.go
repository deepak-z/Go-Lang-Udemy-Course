package main
			 
			 
import "fmt" 
			 
			 
func main() {
			 
	xi := []int{1,2,3};
	w := sum(xi...)
	fmt.Println(w)	
	
	w1 := calculate(sum,xi...)
	fmt.Println(w1)
}

func sum(x ...int) int{
	tot := 0
	for _,v := range x{
		tot += v;
	}
	return tot
}

func calculate(f func(x ...int) int, x ...int) int{
	return f(x...)
}