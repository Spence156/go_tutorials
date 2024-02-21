package main

import (
	"fmt"
	"sync"
	"time"
)

/// Go Routines
/// https://www.youtube.com/watch?v=8uiZC0l4Ajw
// Way of executing code concurrently using parrallel code

var m = sync.RWMutex{} // Can be Mutex which stops all activity or RWMutex which allows reads but not writes while locked
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func dbCall(i int) {
	// Simulate DB Call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	results = append(results, dbData[i])
}

func dbCall1(i int) {
	// Simulate DB Call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock() // Locks the memory object preventing it being written to by other threads
	results = append(results, result)
	m.Unlock() // Unlocks the memory object allowing it being written to by other threads
}

func log() {
	m.RLock() // Unlocks the memory for read access
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock() // Locks the memory object for read access again
}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ { // Using Concurrency but not Parrallel
		dbCall(i)
	}
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Println("\nThe result from the database is:", results)

	results = []string{}
	t0 = time.Now()
	for i := 0; i < len(dbData); i++ { // Using Parrallel
		wg.Add(1)
		go dbCall1(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Println("\nThe result from the database is:", results)
}
