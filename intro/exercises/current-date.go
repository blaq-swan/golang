//usr/bin/true; exec /usr/bin/env go run $0 $@

// Prints the current date

package main

import (
	"fmt"
	"time"
)

func main(){
	currentDate := time.Now()

	fmt.Println(currentDate)
}
