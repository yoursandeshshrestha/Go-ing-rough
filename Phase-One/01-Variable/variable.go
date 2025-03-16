package main

// In Go, fmt is a package from the Go standard library that provides input/output (I/O) functionality. It’s short for format, and it allows you to format and print strings, integers, and other types to the console or standard output, as well as to read input.
import "fmt"

// Variable declarations at package level
var x int
var a, b, c int = 1, 2, 3


var s int = 10000
// This works fine at the package level because you're declaring and initializing the variable in a single step.

// but you cant do this
// var s into
// s = 10000
// This would cause a compilation error if done outside any function (at the package level), because package-level assignments can only be done as part of the declaration itself.

// String declaration
var sandesh string = "sandesh shrestha"
var myfloat float64 = 10.10
var mybool bool = true

const myconst int = 100



func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(s)
	fmt.Println(sandesh)
	fmt.Println(myfloat)
	fmt.Println(mybool)
	fmt.Println(myconst)

	// myconst = 101 // we cannot assign the value to the constant variable
	fmt.Println(myconst)
	// but we can redeclare the variable

	var myconst int = 1011111
	fmt.Println(myconst)

	// first we print the value of a, b, c which is at package level then we assign the value of a, b to 10, 20 and print the value of a, b, if we redeclare the variable a, b then the value get override

    // Variable assignment inside a function
    a = 10
    b:= 20
    d, e, f := 4, 5, "sandesh"

    fmt.Println(a, b)
    fmt.Println(a, b, c)
    fmt.Println(d, e, f)
}

// ======= Extra Infomation =======

// Global (package-level) declarations: You can declare variables outside a function using var. However, you cannot assign values using = or := outside of a function.

// Function-level declarations and assignments: You can use both var and := inside functions, and you can assign values directly.