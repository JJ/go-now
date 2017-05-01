package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/JJ/go-onemax"
)

type eval_chromosome struct {
    chromosome string
    fitness  int
}

// Check out https://gobyexample.com/command-line-arguments
func main() {
	start := time.Now()
	chromosomes := make( chan string )
	evaluated := make( chan eval_chromosome )
	for i := 0; i < 1000; i ++ {
		go func() {
			this_chromosome := random_chromosome( 10000 )
			chromosomes <- this_chromosome
		}()

		go func() {
			this_chromosome := <- chromosomes
			evaluated <- &eval_chromosome{ chromosome: this_chromosome,
				fitness: CountOnes.count( this_chromosome ) }
		}()
	}
	for j := 0; j < 1000; j ++ {
		result := <- evaluated
		fmt.Println( "%10s => %d", result.chromosome, result.fitness )
	}
	fmt.Println("Time ", time.Since(start).Seconds())
}

func random_chromosome( length int ) string {
	chromosome := ""
	for i:= 0; i  < length; i ++  {
		if ( rand.Float32() < 0.5 ) {
			chromosome+= "1"
		} else {
			chromosome+= "0"
		}
	}
	return chromosome
}
