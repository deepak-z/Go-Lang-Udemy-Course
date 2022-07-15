package main

import (
	"fmt" 
	"os"
	"io"
)
			 
			 
func main() {
		
	fmt.Println("Hello, Playground");
	fmt.Fprintln(os.Stdout, "Hello Playground");
	io.WriteString(os.Stdout, "Hello Playground");
			 
}