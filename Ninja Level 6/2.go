package main
			 
			 
import "fmt" 
			 
			 
func main() {
	
	xi := []int{1,2,3,4,5};
	s := sum(xi...)
	fmt.Println(s)

	xii := []int{100,2};
	s1 := sum1(xii)
	fmt.Println(s1)
			 
}

func sum(x ...int) int {
	total := 0;
	for _,v := range x{
		total += v;
	}
	return total
}

func sum1(x []int) int{
	total := 0
	for _,v := range x{
		total += v
	}

	return total
}