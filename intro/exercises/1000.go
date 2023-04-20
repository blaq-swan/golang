//usr/bin/true; exec /usr/bin/env go run $0 $@

// prints 1 through 1000

package main

import "fmt"

func main() {
	for i := 1; i < 1001; i++ {
		fmt.Println(i)
	}
}
