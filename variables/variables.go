package main

import "fmt"

func main(){

	var a = "initial" //automatically infer the type of initialised values
	var b,c int = 1 ,2 // can declare multiple variables at same time
	var truthy = true // can store boolean values
	var e int //automatically the value is zeroed

	f:="newVariable" //shorthand for creating variable

   fmt.Println( a, b, c, truthy, e, f )

}