package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumber() <-chan int {
	randChan := make(chan int)

	go func() {
		defer close(randChan)

		for {
			randomNum := rand.Intn(21) + 10
			randChan <- randomNum
			time.Sleep(time.Second)
		}
	}()

	return randChan
}

func main() {
	randomNumbers := RandomNumber()

	looptime := rand.Intn(5) + 5
	for i := 0; i < looptime; i++ {
		num := <-randomNumbers
		square := num * num
		fmt.Printf("Random number: %d, Square: %d\n", num, square)
	}
}
