//usr/bin/true; exec /usr/bin/env go run $0 $@

// A program that calculates 2 to the power of 11

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Pow(2, 11))
}
