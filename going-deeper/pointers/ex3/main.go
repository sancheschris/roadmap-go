package main

import "fmt"

type Counter struct {
	Count int
}

func (c Counter) Increment() {
	c.Count++
}

func (c *Counter) IncrementPointer() {
	c.Count++
}

func (c *Counter) Decrement() {
	c.Count--
}

func main() {
	c := Counter{0}

	c.Increment() // should NOT change c.Count
	c.Decrement() // should change c.Count

	fmt.Println(c.Count) // Output: -1

	c2 := Counter{0}
	c2.IncrementPointer() // should change c.Count
	fmt.Println(c2.Count) // Output: 1
}