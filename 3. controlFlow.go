package main

import (
	"fmt" 
)
			 
			 
func main() {
	
	//a := 4;
	// for init; condition; post 
	// for i:=0;i<a;i++ {
	// 	fmt.Println(i);
	// }


	/* for i:=0;i<12;i++{
		for j:=0;j<60;j++{
			fmt.Println(i,j);
		}
	} */

	// * for- while statement in go

	/* x := 1;
	for x < 5{
		fmt.Println("Yep");
		x++;
	}


	a := 1
	for{
		if(a > 5){
			break;
		}
		fmt.Println(a);
		a++;
	} */

	// * break and continue

	/* for i := 0; i < 10; i++ {
		if( i == 5) {continue;}
		if(i == 8) {break;}
		fmt.Println(i);

	}
 */


	// * printing ascii

	/* 
	for i := 33; i <= 90; i++ {
		ch := string(i);
		fmt.Println(i,ch);
		fmt.Printf("%v\t%q\n",i,i);
	} 
*/


	// fmt.Printf("%v\t%q\n",200,45);

	// * else if statements

	/* a := 89;

	if(a < 10){
		fmt.Println("Here");
	}else if(a > 10 && a < 50){
		fmt.Println(" in between")
	}else if(a > 50){
		fmt.Println("large");
	}


	if true{
		fmt.Println("oh okay")
	}

	// ? init and condition also goes in if's

	if x := 65; x == 6 {
		fmt.Println("asdf");
	}

	fmt.Println(x) // x is limited to the  */

	/* count := 100
	for i := 0; i < count+1; i++ {
		if(i % 2 != 0) {
			continue;
		}
		fmt.Println(i);
	} */


	// * Switch Statement

	i := 5
	fmt.Println("i is ")
	switch(i){
	case 1 , 5:  // works for both 1 and 5
		fmt.Println("odd");
		//fallthrough // execute the statement below it no matter what
	case 2 : fmt.Println("even");
	case 3 : fmt.Println("yeh bhi odd hain");
	default : fmt.Println("I do not care");
	}

	// ?  a missing switch expression is evaluated to true

	switch {
	case false:
		fmt.Println("this would not print");
	case true:
		fmt.Println("this would print");
	case true:
		fmt.Println("this would also not print");
	}


	//fmt.Println("t && t",true&&true);






}