package main

import "fmt"

func main() {
	//simple if else nesting
	if 10%2 == 0 {        
		fmt.Println("10 is even")
	}else{
		fmt.Println("10 is odd")
	}

	//if without else 
	if 8%2 == 0 {
		fmt.Println("8 is divisble by 2")
	}

	//using logical operators && and ||
	if 10%2 == 0 && 7%2 != 0{
		fmt.Println("10 is even and 7 is odd")
	}

	//conditionals can be nesting with else if and variables declared can be used througout the branching of if else
	if num := 22; num < 0{
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}

//parenthesis are not required around conditionals but braces are required around conditional's body