package main

import (
	"fmt"
)

// func <function_name>(<param1><datatype> , <param2><datatype> ... so on) (<return_type1> , <return_type2> ...so on) {
// 	<function_body>
// }

func add(a int,b int) int {
	return  a + b;
}

func add3(a ,b ,c int) int { // if the parameters are of same type you can omit the datatype for first ones and use only for the last one
	return a + b + c
}

//go has in-built functionality of multiple return values
func sum_diff(a ,b int) (int , int) {
	return a + b , a - b
}

// rest operator and spread operator in golang (used js's terminologies)

// func <><function_name>(<param1> ... <datatype>) (<return_type>){
		// <function_body>
//}

func summation(arr ...int) int {
	sum := 0;
	for _ , val := range arr{
		sum+= val
	}

	return sum
}

func clouser() func()int {
	i := 0

	return func() int {
		i++
		return i
	}
}


func main(){

	fmt.Println(add(1,2))
	fmt.Println(add3(1,2,3))

	sum , diff := sum_diff(1,2)
	fmt.Println(sum , diff)

	_ , ldiff := sum_diff(5345454 ,8645343)
	fmt.Println(ldiff)

	fmt.Println(summation(1,2,3,4,5,6,7,8,9,10))

	arr := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(summation(arr...)) // spreading the array using

	clouserVal := clouser() // clouserVal now holds the reference to the function value that is returned by clouser() and is stored in the memory

	fmt.Println(clouserVal()) // dereference the clouserVal reference to function value
	fmt.Println(clouserVal())
	fmt.Println(clouserVal())

	var fib func(n int) int // declaring an anonymous function's reference in fib variable

	fib = func(n int) int {
		if n < 2{
			return n
		}
		return fib(n -1) + fib(n -2)
	}

	fmt.Println(fib(10))

}