package main

import "fmt"

func main(){

	//single argument for
	i := 0                      
	for i <= 3{
		fmt.Println( "i :", i)
		i = i + 1
	}

   //normal for loop
	for j := 0; j < 3; j++{
		fmt.Println( "j :", j)
	}

	//range - do this N times loop
	for i := range 3 {
		fmt.Println( "range of i :", i)
	}

    // without condition loop until you break out of it or return from the enclosing functions
	for {
		fmt.Println("loop")
		break
	}
	
	//using continue keyword
	for n := range 6 {
		if n%2 == 0{
			continue
		}
		fmt.Println(n)
	}

}