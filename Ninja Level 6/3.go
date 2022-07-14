package main
			 
			 
import "fmt" 
			 
			 
func main() {

	{
		defer after()
		before()
	}
	middle()
			 
}

func before(){
	fmt.Println("this will run first")
}

func after(){
	defer func() {
		fmt.Println("This will run very last defer defer func")
	}()
	fmt.Println("this will run after")
}

func middle(){
	fmt.Println("this will run in middle")
}