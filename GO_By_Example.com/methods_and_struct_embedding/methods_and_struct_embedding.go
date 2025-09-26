package main

import "fmt"

//1️⃣ In go there's no concept of method overloading or overriding like we have in OOP based language . If you want to methods with different signature , there identifier should also be different

//2️⃣ Methods are functions associated with a type .

//3️⃣ Methods are normal functions with a simple twist , if you add a receiver to it then it becomes a method of that type and can only be accesses via that type instance

// 4️⃣ The receiver of a method must be a value or a pointer to a struct type . (not a pointer to an interface or a slice or a map or a channel) .

// 5️⃣ If you have a pointer receiver and call it using a value type or vice-versa , golang compiler will automatically type cast based on the need and get the job done . Cause golang treats pointer and value receiver under the same method set of a type

// 6️⃣ Difference between a value receiver and a pointer receiver is that a value receiver is a copy of the value and a pointer receiver is a reference to the value stored in memory . Thus the value receiver doesn't manipulate what's stored in memory but the pointer receiver does .

// 7️⃣ In general , pointer receiver is more performant than value receiver and is recommended to be used

type shape struct {
	name string
}

func (s *shape) print_name(){
	fmt.Println(s.name)
}

type rect struct{
	shape // embedded struct
	length float32
	breadth float32
}

func (r *rect) area() float32{ // *rect is a pointer to rect struct type
	return r.length * r.breadth
}

func Methods(){

	r := rect{shape: shape{name : "rectangle"} , length : 10 , breadth : 5}
	r.print_name() // embedded structs method is available to us as well
	fmt.Println(r.area())

}