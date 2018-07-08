package lessons

import (
	"fmt"
	"math/rand"
)

func makeRandoms(c chan int) {
	//create a for loop that sends 1000 random numbers down a channel
	for {
		c <- rand.Intn(1000)
	}
}

//BufferedConcurrency  is to show how to send using Buffered Concurrency
func BufferedConcurrency() {
	//create a channel for strings to be sent over
	randoms := make(chan int)

	//start concurrent go routines of emit sending channel as parameter
	go makeRandoms(randoms)

	for n := range randoms {
		fmt.Printf("%d ", n)
	}

}
