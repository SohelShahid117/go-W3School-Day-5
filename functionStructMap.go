// Go Functions

/*
Create a Function

To create (often referred to as declare) a function, do the following:

	Use the func keyword.
	Specify a name for the function, followed by parentheses ().
	Finally, add code that defines what the function should do, inside curly braces {}.
*/
package main

import "fmt"

// Create a function
func myMessage() {
	fmt.Println("Hello World!")
	fmt.Println("I just got executed!")
}

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func familyName2(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}
func myFunction(x int, y int) int {
	return x + y
}

func myFunction2(x int, y int) (result int) {
	result = x + y
	return result
}

func myFunction3(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction4(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func myFunction5(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
func myFunction6(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
func myFunction7(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	myMessage() // call the function
	fmt.Println("function call again")
	myMessage() // call the function

	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
	/*
		Note: When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.
	*/

	familyName2("Liam", 3)
	familyName2("Jenny", 14)
	familyName2("Anja", 30)
	/*
		Note: When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, and the arguments must be passed in the same order.
	*/

	fmt.Println(myFunction(1, 2))

	fmt.Println(myFunction2(11, 22))

	total := myFunction3(11, 2)
	fmt.Println(total)

	fmt.Println(myFunction4(15, "Hello"))

	a, b := myFunction5(5, "Hello")
	fmt.Println(a, b)

	_, y := myFunction6(15, "Helloooo")
	fmt.Println(y)

	p, _ := myFunction7(51, "Hello")
	fmt.Println(p)

}
