package main

import (
	"fmt"
)

func byVal(a int){
	a = 0	// this does nothing s the primitives are passed by copy thus this function body manipulations doesn't affect the original value of a
}

func byRef (a *int){
	*a = 0 // dereferencing the passed address and then changing the value at the address
}


func main(){

	a := 10

	byVal(a)
	fmt.Println(a)

	byRef(&a)
	fmt.Println(&a) // &a is the address of a's value in the memory
	fmt.Println(a)

}

