package main
			 
			 
import "fmt" 
			 
			 
func main() {
	
	for i:='A';i<='Z';i++{
		for j:=1;j<=3;j++{
		fmt.Printf("%v\t%#U\n",i,i);
		}
	}
			 
}