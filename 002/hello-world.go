// required by every program
package main

// package import
import "fmt"

// entry point into program
func main() {
	// print Hello whole world
	fmt.Println("Hello whole world.")
	// call foo function
	foo()
	// print Something else
	fmt.Println("Something else")

	// start at 0
	// loop through numbers less than 100
	// increment by one each iteration
	// and if i is even
	// print i
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// call bar function to log exit
	bar()
}

func foo() {
	// print I'm in foo
	fmt.Println("I'm in foo")
}

func bar() {
	// print ..and then exit
	fmt.Println("..and then exit")
}
