package main
			 
			 
import "fmt" 

type person struct{
	first string
	last string
	ficf []string
}
			 
			 
func main() {
		
	p1 := person{
		first: "Deepak",
		last: "Kumar",
		ficf: []string{"pista","butterscotch"},
	}

	p2 := person{
		first: "Deepak",
		last: "Kumar from future",
		ficf: []string{"pista","butterscotch"},
	}

	/* fmt.Println(p1)
	fmt.Println(p2)
 */
/* 	for i,v := range p1.ficf{
		fmt.Println(i,v)
	} */

	mp := map[string]person{
		p1.last : p1,
		p2.last : p2,
	}

	/* fmt.Println(mp) */

	for k,val := range mp{
		fmt.Println(k)
		for i,v := range val.ficf{
			fmt.Println(i,v)
		}
	}


			 
}