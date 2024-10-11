package main

import (
	"fmt"
	"time"
)

func main() {
	//basic switch
	i := 4
	fmt.Print("Write ", i, " as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("either i < 0 or i > 3 ")
	}

	//default case is optional

	//we can comma separate multiple expressions in cases
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's weekday")
	}

	//switch without expression is an alternate way to express if/else logic, here we can show that the case expressions can be non constants
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println(t.Hour())
		fmt.Println("it's before noon")
	default:
		fmt.Println(t.Hour())
		fmt.Println("it's after noon")
	}

	//we can also implement type switch, which compares types instead of values

	whatAmI := func (i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println("is boolean")
		case int:
			fmt.Println("is integer")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}

	whatAmI(2)
	whatAmI(false)
	whatAmI(2.5)

}