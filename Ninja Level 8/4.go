package main
			 
			 
import (
	"fmt"
	"sort"
) 
			 
			 
func main() {
		
	xi := []int{5,2,3}
	xs := []string{"asdf","asdfasdf","zzz","l"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)

			 
}