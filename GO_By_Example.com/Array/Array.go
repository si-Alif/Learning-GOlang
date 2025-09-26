package main

import (
	"fmt"
)

func main(){
	// structure : var <identifier> [array_length]<type> = <value>
	var arr [5]int  = [5]int{3 ,5,6} // for preassigning values make sure that the array length you're trying to assign is the same length as the array length of the declared array length

	fmt.Println(arr) // --> [3 , 5 , 6 , 0 , 0] --> as the array length is 5 but ony first 3 elements were ever declared


	// declaring an array without preassigning values will result in zero value of the type
	var arr2 [5]int
	fmt.Println(arr2) // --> [0 , 0 , 0 , 0 , 0]

	// length of an array
	fmt.Println(len((arr)))

	// local variable scope declaration
	local_arr := [5]string{"John Doe" , "jane Doe"  , "John Doe" , "jane Doe" , "John Doe" }

	fmt.Println(local_arr[3])

	// let the compiler figure out the length of the array
	len_arr := [...]string{"unga" , 3:"bunga" , "zanga"} // here in the assigned value , we did <index> : <value> which compiler reads as place <value> at this <index> on the array

	// here in the output , you can observe that index 1 and 2's value is zero("")
	for i , val := range len_arr{
		fmt.Println("Index : " , i  , "value : " , val)
	}

	var twoDim[2][3] int = [2][3]int{
									{2, 3 ,7},
									{5 ,8, 1},
	}

	for i:= range 2{
		for j:= range 3{
			fmt.Println("Element " , i ,"x" , j , "=" , twoDim[i][j])
		}
	}

	




}