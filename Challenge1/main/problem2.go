package main

import (
	"log"
    "sync"
	"time"
	"math/rand"
)

// declare a variable that holds the go routines queue sync execution
var wg1 sync.WaitGroup

// declare a var to lock the num var while it will increment its value
var mutex1 = &sync.Mutex{}

func problem2() {

	log.Printf("problem2: started --------------------------------------------")

	//
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//
	
	for inx := 0; inx < 10; inx++ {
		
		// adding the next go routine to the wait group queue
		wg1.Add(1)

		go printRandom2(inx)
		
	}

	// instruction that allows problem1 function to wait for the ending of the go routines executions
	wg1.Wait()
	
	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int) {
	for inx := 0; inx < 10; inx++ {
		
		// locking the printing of rand number...
		mutex1.Lock()

		// sleeps for a second
		time.Sleep(time.Duration(1000) * time.Millisecond)

		// then print the random number
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())

		// unlocking the printing of rand number
		mutex1.Unlock()

	}

	// says that this routine is ended
	wg1.Done()
}