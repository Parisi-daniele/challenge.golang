package main

import (
	"log"
    "sync"
	"os"
	"math/rand"
)

// declare a global var that counts the total numbers to be printed
var num = 0

// declare a var to lock the num var while it will increment its value
var mutex = &sync.Mutex{}

// declare a variable that holds the go routines queue sync execution
var wg sync.WaitGroup

func problem1() {

	log.Printf("problem1: started --------------------------------------------")

	//
	// Todo:
	//
	// Quit all go routines after
	// a total of exactly 100 random
	// numbers have been printed.
	//
	// Do not change the 25 in loop!
	//


	for inx := 0; inx < 10; inx++ {

		// adding the next go routine to the wait group queue
		wg.Add(1)

		go printRandom1(inx)

	}

	// instruction that allows problem1 function to wait for the ending of the go routines executions
	wg.Wait()

	log.Printf("problem1: finised --------------------------------------------")
}

func printRandom1(slot int) {

	//
	// Do not change 25 into 10!
	//
	defer wg.Done()
	for inx := 0; inx < 25; inx++ {

		// locking the num var...
		mutex.Lock()
		
		num++ // increment the num var value

		// unlocking the num var
		mutex.Unlock()

		// if the num var is equal to 100...
		if num == 100 {
			// exit from the go routine
			os.Exit(0)
		}else{
			// otherwise print the rand number
			log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		}

	}
}