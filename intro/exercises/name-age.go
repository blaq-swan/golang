//usr/bin/true; exec /usr/bin/env go run $0 $@

// A program to print name and age

package main

import "fmt"

func main() {
	first_name := "Blaq"
	var last_name string = "Swan"
	var age int = 45

	fmt.Printf("My name is %v %s and I am %d years old.\n", first_name, last_name, age)
}
