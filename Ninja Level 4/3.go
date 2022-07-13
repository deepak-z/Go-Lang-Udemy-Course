package main

import "fmt"

func main() {

	x := []int{1, 2, 7, 3, 4, 5, 6, 8, 9, 10}

	fmt.Println(x[2:6]);
	fmt.Println(x[2:]);
	fmt.Println(x[:3]);
}
