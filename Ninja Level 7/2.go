package main
			 
			 
import "fmt" 
	
type person struct{
  name string
}

/* func (p *person) changeName(){  
//func (p person) changeName(){ -> this would not  change the value
	p.name = "HaHaha"
} */


func main() {
			 
		p := person{
			name: "Deepak",
		}	
		
		
		fmt.Println(p)
		changeNamee(&p)
		fmt.Println(p)
}

func changeNamee(p *person){
	//p.name = "Changed in function"
	(*p).name = "Hey"
}