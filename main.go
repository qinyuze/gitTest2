package main

import "fmt"

func main() {
	fmt.Println("hello 1")
	fmt.Println("hello 2")

	// after first commit
	fmt.Println("hello 3")

	// master commit
	fmt.Println("hello master")

	// change in hot-fix
	fmt.Println("hello hot-fix")

	// after push to github
	fmt.Println("hello push to github")
	
	// test pull
	fmt.Println("hello test pull")
}
