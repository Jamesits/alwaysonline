package main

import "log"

var softErrorCount uint64

// if QuitOnError is true, then panic;
// else go on
func softFailIf(e error) {
	if e != nil {
		softErrorCount++
		log.Printf("[ERROR] %s", e)
	}
}

func hardFailIf(e error) {
	if e != nil {
		panic(e)
	}
}