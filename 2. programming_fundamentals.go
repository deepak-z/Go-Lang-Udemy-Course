package main

import (
	"fmt" 
	//"runtime"
)
			 

var x bool
var b float64
var xx uint8

func main() {
	
	// xx = 154;
	// b = 9.88;
	// fmt.Println(b);
	// fmt.Println(x);
	// fmt.Printf("%T\n",x);

	// a := 7;
	// b := 42;

	// fmt.Println(a == b);
	
	// a := 42
	// b := 43.22
	// fmt.Printf("%T\n",a);
	// fmt.Println(b);
	// fmt.Printf("%T\n",b);

	// fmt.Println(runtime.GOOS);
	// fmt.Println(runtime.GOARCH);

	// STRINGS

/* 	s := "Hello, Playground"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	bs := []byte(s);  // slice of bytes
	fmt.Println(bs)
	fmt.Printf("%T\n",bs); */


	//  CONSTANTS

	/* const a = 65;

	fmt.Println("sdf");
	fmt.Println(a);

	//a = 89;

	const(
		b = 45.454
		d = "asdfasdf"
	);

	fmt.Println(b);
	fmt.Println(d); */


	// IOTA

	// if you need something to automatically increment it by one
	// one can use iota

	/* const (
		a = iota; // a = 4;
		b;
		c;
	)

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c); */


	// BINARY SHIFTING

	a := 4;
	fmt.Println(a>>1);
	fmt.Println(a << 1);

	const (
		_ = iota;
		kb = 1 << (iota * 10);  // 1 * 10
		mb = 1 << (iota * 10);  // 2 * 10
		gb = 1 << (iota * 10);  // 3 * 10
	)

	fmt.Println(kb);
	fmt.Println(mb);
	fmt.Println(gb);

			 
}