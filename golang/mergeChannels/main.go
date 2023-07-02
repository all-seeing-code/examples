package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getChan(vs ...int) chan int {
	ch := make(chan int)
	go func() {
		for _, v := range vs {
			ch <- v
			time.Sleep(time.Duration(rand.Intn(1000) * int(time.Millisecond)))
		}
		close(ch)
	}()

	return ch
}

func mergeChannels(a, b chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for a != nil || b != nil {
			select {
			case t, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				ch <- t
			case t, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				ch <- t
			}
		}
	}()

	return ch
}

func main() {

	a := getChan(1, 2, 3, 4, 5, 6, 7, 8, 9)
	b := getChan(11, 12, 13, 14, 15, 16, 17, 18, 19)

	d := mergeChannels(a, b)

	for v := range d {
		fmt.Println(v)
	}

	time.Sleep(100 * time.Millisecond)
}
