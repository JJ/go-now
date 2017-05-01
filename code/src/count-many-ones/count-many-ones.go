package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/JJ/go-onemax"
	"sync"
)

type eval_chromosome struct {
    chromosome string
    fitness  int
}


// Check out https://gobyexample.com/command-line-arguments
var wg sync.WaitGroup 
func main() {
	n := 1000
	start := time.Now()
	chromosomes := make( chan string, n )
	evaluated := make( chan *eval_chromosome, n )
	for i := 0; i < n; i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			this_chromosome := random_chromosome( 10000 )
			chromosomes <- this_chromosome
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			this_chromosome := <- chromosomes
			this_evaluated := &eval_chromosome{ chromosome: this_chromosome,
				fitness: CountOnes.Count( this_chromosome ) }
			evaluated <- this_evaluated
		}()
	}

	wg.Wait()
	i:= 1
	close(evaluated)
	for result := range evaluated {
		fmt.Printf("%d => %d\n", i, result.fitness )
		i++

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
