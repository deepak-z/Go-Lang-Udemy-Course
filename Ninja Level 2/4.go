package main

import "fmt"
			 
func main() {
	
	var a int;
	a = 4;
	fmt.Println(a);
	fmt.Printf("%d\t%b\t%#x\n",a,a,a);
	aa := a >> 1;
	fmt.Println(aa);
	fmt.Printf("%d\t%b\t%#x\n",aa,aa,a);
			 
}