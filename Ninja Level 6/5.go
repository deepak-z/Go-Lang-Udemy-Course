package main
			 
			 
import "fmt" 
	
type sqaure struct{
	side int
}

type circle struct{
	radius int
}

func (s sqaure) area() int{
	return s.side * s.side
}

func (c circle) area() int{
	return 3 * c.radius * c.radius
}

type shape interface{
	area() int
}

func printShape(sh shape){
	fmt.Println("Area now : ",sh.area())
}


func main() {
	
	sq := sqaure{
		side : 4,
	}

	cr := circle{
		radius: 4,
	}

	fmt.Println(sq.area())
	fmt.Println(cr.area())

	printShape(cr)
	printShape(sq)
			 
}