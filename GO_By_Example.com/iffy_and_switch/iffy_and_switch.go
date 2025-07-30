package main

import (
	"fmt"
	"time"
)

func main(){
	curr_day := time.Now().Weekday()
	switch curr_day {
	case time.Friday , time.Saturday :
		fmt.Println("It's the weekend!")
	default :
		fmt.Println("It's a weekday.")
	}

	// Example of a switch statement with no condition
	// instead of using if-else, we can use switch to check multiple conditions
	switch {
		case time.Now().Hour() < 12:
			fmt.Println("Good Morning!")
		case time.Now().Hour() < 18:
			fmt.Println("Good Afternoon!")
		default:
			fmt.Println("Good Evening!")
	}

	
}


