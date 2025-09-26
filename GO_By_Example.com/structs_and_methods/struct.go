package main

import(
	"fmt"
)

// struct in go is pretty similar to struct in C++ (almost)
// main purpose of struct is to group related data together

type person struct {
	name string
	age int
}

func createPerson(name string) *person{
	p := person{name : name} // instantiate / construct the person
	p.age = 21 // set age using the dot operator
	return &p // return address of the newly created person struct
}


func main(){
	fmt.Println(person{name:  "bob" , age: 20}) // using data field names
	fmt.Println(person{"alice" , 30}) // using place holder
	fmt.Println(person{name : "max"}) // if a filed is not initialized , it will be set to zero value

	fmt.Println(&person{name : "ann" , age : 40}) //a pointer to the address where the struct is stored in the memory

	p := person{"john" , 50}
	fmt.Println(p.name) // accessing fields using dot operator

	ptr := &p
	fmt.Println(ptr.name) // "john" --> golang automatically dereferences and gets the value using the dot operator .

	ptr.age = 55
	fmt.Println(ptr.age) // structs are mutable as there's no concept of encapsulation in golang

	fmt.Println(createPerson("alif"))

	// creating an anonymous struct and defining it . Anonymous structs are useful when you'll use it for just once
	anonymous_struct := struct{
		name string
		age int
		birthday string

	}{
		name : "John Doe",
		age : 25,
		birthday : "01-01-2000",
	}

	fmt.Println(anonymous_struct)

}