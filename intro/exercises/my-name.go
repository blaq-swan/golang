//usr/bin/true; exec /usr/bin/env go run $0 $@

// A program that prints my name

package main

import "fmt"

func main(){
	var first_name string = "Blaq"
	var last_name string = "Swan"

	fmt.Printf("My name is %s %s\n", first_name, last_name)
	fmt.Printf("My name is %v %v\n", first_name, last_name)
	fmt.Println("My name is", first_name, last_name)
}
