package main

import "fmt"

func main(){
	i:= 1
	for i <=10{
		fmt.Println("Iteration:", i)
		i += 4
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0{
			fmt.Println("Even Number:", i)
		}else {
			fmt.Println("Odd Number:", i)
		}
	}

	for {
		fmt.Println("Infinite Loop")
		break // to prevent infinite iterations
	}

	arr := []string{"a" , "b" , "c" }

	for i , val := range arr{ // val is copy of the element not the reference
		fmt.Println("Index : " , i  , "value : " , val)
	}

	str := "Hello, World!"
	for i , char := range str{
		fmt.Printf("Index: %d, Character: %c\n", i, char) // placeholders used to print values
	}
	

}