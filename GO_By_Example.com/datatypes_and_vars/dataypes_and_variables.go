package main // entrypoint package
// every package can be compiled into a binary code but they won't be executable


import (
	"fmt"
)

// GO is heavily influenced unix like structure and simplicity
// main package is the root package of the entire application and under this you can have multiple packages which might contain multiple files
// This language doesn't have any export / public / private keywords
// Instead, it uses capitalization to determine visibility of identifiers(variables / functions / type) across packages
// Identifier --> starts with an uppercase letter --> it's exported (public).
// Identifier --> starts with a lowercase letter --> it's unexported (private to its package).

const int_val_max int = 10e9

func main(){ // main function + main package combined is the entrypoint to the entire GO application , when compiled you can use double click / CLI based system to execute that binary code

	fmt.Println("Hello, World!") // reason why "P" is capitalized is because of the simplicity of exporting identifiers across packages
	fmt.Println("This is a simple Go program.")

	fmt.Println("go" + "lang"); // string concatenation
	fmt.Println("3.45 / 2.31 = " ,  3.45 / 2.31) // float64 is default to store decimal numbers in GO
	fmt.Println(true && false) // AND
  fmt.Println(true || false) // OR
  fmt.Println(!true) // Negation

	// variable declaration structure(both local and package level) ---> var <identifier> <type> = <value> (if any value isn't assigned then it's defaulted to zero value of the type)
	var price float32 = 19.49 // --------> var price float32 = 19.49
	fmt.Println("Price:", price)

	// variable declaration structure(only at functional scope) ---> <identifier> := <value>
	quantity := 5 // --------> quantity := 5
	fmt.Println("Quantity:", quantity)

	// constant declaration structure(think of it like #define from C but more precise about type and scope ; can be local or universal) ---> const <identifier> <type> = <value>
	const pi float64 = 3.14159 // --------> const pi float64 = 3.14159
	// constant values in Go are immutable and can be typed or untyped


	// iota
	// iota is a pre-declared identifier in Go that simplifies the creation of enumerated constants (from google)
	// Think of it as Go's minimal, built-in way to define enums or flags - without needing an enum keyword.
	// iota is reset to 0 whenever the keyword const appears in the source code, and it increments by 1 for each constant declaration within that const block.
	const (
		Friday = iota // 0
		Saturday      // 1
		Sunday        // 2
		Monday  			// 3
		Tuesday       // 4
		Wednesday     // 2
		Thursday      // 6
	)

	// real world example of iota
	const (
		_ = iota // ignore first value by assigning to blank identifier (yeah , "_" is a blank identifier in Go )
		statusOK // 1
		statusCreated = 201 // should have been 2 but we are explicitly setting it to 201
		statusAccepted = iota // resumes with 3
		statusBadRequest = 0
		statusNoContent = 0
		statusInternalServerError = 0
		statusNotImplemented = 0
		statusServiceUnavailable = 0
	)

}



