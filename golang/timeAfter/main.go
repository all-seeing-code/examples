package main

import (
	"fmt"
	"runtime"
	"time"
)

func handle(event string) {
	fmt.Println("Received: ", event)
}

// func consumer(ch <-chan string) {
// 	for {
// 		select {
// 		case event := <-ch:
// 			handle(event)
// 		case <-time.After(100 * time.Millisecond):
// 			// Because of the surrounding for loop, this timer gets reset
// 			// on each iteration on receiving successful event
// 			log.Println("warning: no messages received")
// 		}
// 		time.Sleep(20 * time.Millisecond)
// 	}
// }

func consumerWithTimer(ch <-chan string) {
	timer := time.NewTimer(100 * time.Millisecond)
	for {
		timer.Reset(2000 * time.Millisecond)
		select {
		case <-ch:
			// handle(event)
		case <-timer.C:
			// log.Println("warning: no messages received")
		}
		// time.Sleep(20 * time.Millisecond)
	}
}

// receive only channel
func consumer(ch <-chan string) {
	for {
		select {
		case <-ch:
			// fmt.Println("Received: ", event)
		case <-time.After(1 * time.Hour):
			// fmt.Println("warning: no messages received")
		}
		// time.Sleep(2 * time.Millisecond)
	}

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func getMemory() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Println("Memory Usage:")
	fmt.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(memStats.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(memStats.Sys))
	fmt.Printf("\tNumGC = %v\n", memStats.NumGC)
}

func simulateTimeAfter() {
	start := time.Now()

	// Simulating multiple requests
	for i := 0; i < 1000000; i++ {
		select {
		case <-time.After(1 * time.Hour):
			// Simulating a long timeout that might not be triggered soon.
		default:
			// Simulating request processing.
		}
	}

	// Check the memory statistics after the loop
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Println("Memory Usage:")
	fmt.Printf("Alloc = %v MiB", bToMb(memStats.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(memStats.TotalAlloc))
	fmt.Printf("\tSys = %v MiB\n", bToMb(memStats.Sys))
	fmt.Printf("Execution Time: %v\n", time.Since(start))
}

func main() {

	// simulateTimeAfter()
	ch := make(chan string)

	go consumerWithTimer(ch)

	events := []string{"one", "two", "three"}
	count := 0
	for {
		ch <- events[0]
		// time.Sleep(10 * time.Millisecond)
		ch <- events[1]
		// time.Sleep(10 * time.Millisecond)
		ch <- events[2]
		count++
		if count%10000 == 0 {
			getMemory()
		}
	}

}
