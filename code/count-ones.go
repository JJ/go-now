package main

import (
	"os"
	"fmt"
	"time"
)

// Check out https://gobyexample.com/command-line-arguments
func main() {
	start := time.Now()
	unos := 0
	for _, element := range os.Args[1] {
		if ( element == '1' ) {
			unos++
		} 
	}	
	fmt.Println("Unos ,", unos,", ", time.Since(start).Seconds())
}
