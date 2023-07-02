package main

import (
	"fmt"
	"sync"
)

type Customer struct {
	mu  sync.RWMutex
	id  string
	age int
}

func (c *Customer) UpdateAge(age int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if age < 0 {
		return fmt.Errorf("age should be positive for customer %+v", c)
	}
	c.age = age
	return nil
}

func (c *Customer) String() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return fmt.Sprintf("id %s, age %d", c.id, c.age)
}

func main() {
	c := &Customer{id: "1", age: 15}
	err := c.UpdateAge(-10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c.age)
	}
}
