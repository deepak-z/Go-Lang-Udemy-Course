package main
			 
			 
import "fmt" 
			 
			 
func main() {
			 
	as := struct{
		do_I_Care bool
	}{
		do_I_Care: false,
	} 

	fmt.Println(as.do_I_Care)
}