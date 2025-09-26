package main

import (
	"fmt"
	"slices"
	// "slices"
)

func main(){
	var s []string ; // as the slice is uninitialized it's value is nil , means it points to nothing
	fmt.Println("Before :" , s , s== nil , len(s) == 0)

	s = make([]string, 3) // initializes the slice elements with zero values
	fmt.Println("After :" , s , s== nil , len(s) == 0) // elements are now ""

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// s[3] = "d" // panic: runtime error: slice bounds out of range [3] with capacity 3 , as it's length is 3 as we used make to initialize the slice
	fmt.Println("set :" , s)

	s = append(s, "d") // this works fine as append dynamically increases the size of the slice and assigns the value
	s = append(s, "e" , "f" , "g") // appending multiple values at once
	fmt.Println("set :" , s)

	copySlice := make([]string , len(s))
	copy(copySlice , s) // here copySlice is now pointing to a completely different value which is equal to value of "s"

	sliced := s[2 : 5] // here sliced shares array with s , means they refers to the same array value in the memory . Thus , manipulating sliced will also affect s

	fmt.Println(s[2]) // "c"
	sliced[0] = "nicholas jackson, the goat"
	fmt.Println(s[2]) // "nicholas jackson, the goat"

	latter_part := s[2:] // these kind of slicing is same as we do in python where <slice>[<low> : <high>] ---> low is inclusive and high is exclusive
	fmt.Println(latter_part)

	a := []string{"a" , "b" , "c" , "d" , "e" , "f" , "g"}
	b := []string{"a" , "b" , "c" , "d" , "e" , "f" , "g"}
	fmt.Println(slices.Equal(a , b))

	

}