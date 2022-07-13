package main
			 
			 
import "fmt" 
			 
			 
func main() {
			 
	x := [][]string{{"I"},{"actually"},{"do"},{"not"},{"care"}};
	
	for i := range x{
		for j := range x[i]{
			fmt.Println(x[i][j])
		}
	}
}