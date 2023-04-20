//usr/bin/true; exec /usr/bin/env go run $0 $@

// Generate a random number between 0 and 10

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(rand.Intn(10-0) + 0)
}
