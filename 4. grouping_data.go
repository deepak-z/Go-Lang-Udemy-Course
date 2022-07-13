package main
			 
			 
import "fmt" 
			 
			 
func main() {
		
	/* var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x)) */

	// * SLICES -> allows you to group together values of the same type
	// composite literals
	//x := []int{1,2,3,4,5};
	/* fmt.Println(x);
	fmt.Println(len(x));
	fmt.Println(cap(x));

	fmt.Println(x[1]); */
	
	// range printing of slice values

	/* for i,v := range x{
		fmt.Println(i,v);
	}*/
	
	/*

	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[1:4])
 */
	/* for i := 0; i < len(x); i++ {
		fmt.Println(i)
	} */

	// * APPENDING TO A SLICE

	/* fmt.Println(x);
	x = append(x, 4,5)
	fmt.Println(x);

	y := []int{12,13,14}
	//z := []string{}
	y = append(y, x...);  // variadic
	fmt.Println(y) */

	// * DELETING FROM A SLICE

	/* x = append(x[:1],x[3:]...)
	fmt.Println(x) */


	// * slice - make

	/* x := make([]int, 10,100);
	fmt.Println(x);
	fmt.Println(len(x));
	fmt.Println(cap(x)); */

	// * Multi dimensional - slice

	/* jb := [][]string{{"hey","I","don't","care","once"},{"hey","I","don't","care","twice"}};
	fmt.Println(jb); */

	// * MAPS

	mp := make(map[string]int);
	mp["asdf"] = 1
	mp["asdfasdf"] = 2

	fmt.Println(mp);


	mp1 := map[string]int{
		"y":1,
		"z":2,
	}

	fmt.Println(mp1);
	fmt.Println(mp1["y"]);
	fmt.Println(mp1["yf"]); // this will by default return 0 if something does not exist
	
	v,ok := mp1["yf"]
	fmt.Println(v)
	fmt.Println(ok) // this would be false

	if v,ok := mp1["yf"]; ok {
		fmt.Println("This would not be printed",v);
	}

	mp1["yf"] = 56

	fmt.Println(mp1["yf"]);

	for str,value := range mp1{
		fmt.Println(str,value);
	}


	delete(mp1,"yf");

	for str,value := range mp1{
		fmt.Println(str,value);
	}

	if v,ok := mp1["yff"]; ok{
		delete(mp1,"yff");
		fmt.Println("this would not get printed",v);
	}






			 
}