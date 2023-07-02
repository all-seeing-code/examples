package main

import "fmt"

type Vehicle struct{}

func (v Vehicle) honk() {
	fmt.Printf("honk honk\n")
}

type Car struct {
	Name string
	Vehicle
}

func main() {
	c := Car{"Mustang", Vehicle{}}
	c.honk()
}
