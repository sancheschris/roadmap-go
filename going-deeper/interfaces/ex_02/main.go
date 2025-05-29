package main

import (
	"fmt"
)

/*
Interfaces are implicit" means:
You don’t declare anywhere that your struct implements an interface.

In Java, TypeScript, C#, and others, you have to explicitly say:
Person implements Greeter

In Go, that’s unnecessary — if the method matches, it works. That's structural typing

The compiler checks method signatures to determine compatibility.
*/

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

// Person "implements" Greeter just by having a method with the same signature
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	p := Person{Name: "Alice"}
	sayHello(p) // works, even though we never wrote "implements"
}