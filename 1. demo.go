package main

import "fmt"

func main(){
	//1. short decleration operator
	// x := 42
	// y := 100 + 24
	// name := "Deepak Kumar"
	// fmt.Println(x,y,name);
	// fmt.Println("Hello world");

	//2. Var keyboard

	// var y = 43;
	// fmt.Println(y);

	//3. format printing

	// fmt.Printf("%#x\t%b\t%x\n",y,y,y);
	// s := fmt.Sprintf("%#x\t%b\t%x",y,y,y);
	// fmt.Println(s)

	//4. creating your own type - GO IS ALL ABOUT TYPES

	var a int;
	a = 42;
	fmt.Println(a);
	fmt.Printf("%T\n",a);

	type hotdog int;
	var b hotdog; 
	b = 43;
	fmt.Println(b);
	fmt.Printf("%T\n",b);

	// a = b; // both are not same

	//5. to convert one to another we can do the following

	// this is called conversion

	a = int(b);






	
}