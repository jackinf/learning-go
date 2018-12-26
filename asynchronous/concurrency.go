package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func runConcurrency() {
	wg.Add(2)
	go printCounts("label 1")
	go printCounts("label 2")
	wg.Wait()
	fmt.Println("Done")
}

func printCounts(label string) {
	defer wg.Done()
	for count := 0; count < 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count %d from label %s\n", count, label)
	}
}