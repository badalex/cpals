package main

import (
	"cpals/mt"
	//	"fmt"
	"os"
	"runtime"
	"time"
)

var maxSleep uint = 1000

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	sleepM := mt.MT{}

	// wait a random number of seconds
	sleepM.Seed(uint(time.Now().Unix() ^ int64(os.Getpid())))
	s := time.Duration(sleepM.Rand()%maxSleep) * time.Second
	//fmt.Printf("pretend sleeping: %d\n", s/time.Second)

	// seeds the rng with the current unix timestamp
	seed := uint(time.Now().Unix()) + uint(s/time.Second)
	//fmt.Printf("seed is %d\n", seed)
	m := mt.MT{}
	m.Seed(seed)

	// wait a random number of seconds
	s = time.Duration(sleepM.Rand()%maxSleep) * time.Second
	s = 0
	//fmt.Printf("pretend sleeping: %d\n", s/time.Second)

	// get the first 32 bits
	r := m.Rand()

	// figure out the seed we now that it had to been seeded in the last couple
	// of minute so just brute it
	nunix := seed + uint(s/time.Second) + 1
	nseed := nunix - (maxSleep)

	prog := make(chan uint)
	got := make(chan uint)
	for i := nseed; i < nunix; i++ {
		go func(n uint) {
			m1 := mt.MT{}
			m1.Seed(n)
			rr := m1.Rand()
			//fmt.Printf("trying %d %d %d\n", n, nunix, seed)

			if rr == r {
				got <- n
				return
			}
			prog <- 1
		}(i)
	}

	timeout := time.After(30 * time.Second)

	for {
		select {
		//case gseed := <-got:
		case _ = <-got:
			//fmt.Printf("found seed %d\n", gseed)
			return
		case _ = <-timeout:
			panic("failed to find seed, timedout")
		case _ = <-prog:
			//fmt.Printf(".\n")
		}
	}
}
