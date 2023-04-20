//usr/bin/true; exec /usr/bin/env go run $0 $@

// A program that prints numbers 1 through 10

package main

import "fmt"

func main(){
	for i := 0; i <= 10; i++ {
		fmt.Printf("%v ", i)
	}
}
