package main
			 
			 
import "fmt" 

type person struct{
	first string
	last string
}

type secretAgent struct{
	person
	ltk bool
}

// * method

// ? according to the receiver the function gets attached to the particular struct

func ( s secretAgent) tellName(){
	fmt.Println("I am ",s.first,s.last,"-the secret Agent");
}

func ( p person) tellName(){
	fmt.Println("I am ",p.first,p.last,"-the person");
}



// * Interfaces and Polymorphism

// ! Keyword identifier type

// ? now secretAgent and person are now of both type human
type human interface{
	tellName()

	// ? interface says if you have this method then you are my type
}

func printHuman(h human){
	fmt.Println("I called human",h)
}


// * NOTE : printLn has a empty interface , so that means that every other function which does not have
// * any function , can be used in println , which we do 

// todo : ASSERTION CODE 


			 
			 
func main() {

	   // ! func ( r receiver) identifier(parameters) (return(s)) { code }
			 
		/* foo()	
		bar("Deepak") 
		s1 := woo("Deepak")
		fmt.Println(s1)
		s2,s3 := fizzbuzz("Ninja","Hattori")
		fmt.Println(s2)
		fmt.Println(s3) */


		// * Unlimited parameters - defurling

		/* foo1(2,3,4,5,6,7,8);

		x1 := []int{1,2,3}

		foo1(x1...)
		foo1() */

		// * Defer

		/* defer de(); // this would now run after containgin function [ func main ()  in this case] 
		// reaches '}' its closing tag or the end
		fer(); */

		// * Methods

		sa1 := secretAgent{
			person: person{
				first: "Deepak",
				last: "Kumar",
			},
			ltk: true,
		}

		/*

		fmt.Println(sa1)
		sa1.tellName(); */

		// * Interfaces and polymorphism

		p1 := person{
			first: "IDC",
			last: "Kumar",
		}

		// this can work with both person and secret Agent because
		// secretAgent and person are both of type human
		printHuman(p1)
		printHuman(sa1)




}


func de(){
	fmt.Println("deeeee.....")
}

func fer(){
	fmt.Println("ferrrr..........")
}
func foo1(x ...int){ // the variaditic parameter has to be the last para meter
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	sum := 0;
	for i,_ := range x{
		sum += x[i];
	}

	fmt.Println(sum)
}

func foo(){
	fmt.Println("This has been printed in the function")
}

func bar(s string){
	fmt.Println("Hello",s);
}

func woo(s string) string {
	s = s + "I don't care"
	return s
}

func fizzbuzz(fn string,ln string) (string,bool){
	return "hello",true
}