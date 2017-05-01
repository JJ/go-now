package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var unos_str string = os.Args[1]
	chromosome := make([]bool,len(unos_str))
	for index, _ := range chromosome {
		if ( unos_str[index] == '1' ) {
			chromosome[index] = true
		} else {
			chromosome[index] = false
		}
	}
	
	unos := 0
	element := false
	for len(chromosome)>0 {
		element, chromosome = chromosome[0], chromosome[1:]
		if ( element ) {
			unos++
		}
	}
	fmt.Println("Unos ,", unos,", ", time.Since(start).Seconds())
}
