package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println("f", from, ":", i)
	}
	fmt.Println("End f ---")
}

func func2(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println("func2", from, ":", i)
	}
	fmt.Println("End func2! ****")
}

func mainHello() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")

	go func2("goroutine inside main")

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.
	fmt.Println("Press any key to END")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Bye :)")
}
