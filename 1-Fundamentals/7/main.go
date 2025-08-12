package main

import "fmt"

func main() {

	salary := map[string]int{"Christian": 1000, "John": 2000, "Mary": 3000}
	fmt.Println(salary["Christian"])
	delete(salary, "Mary")
	salary["Gabi"] = 5000
	fmt.Println(salary["Gabi"])
	fmt.Println(salary)
	salary["Adalberto"] = 15000
	fmt.Println(salary["Adalberto"])
	fmt.Println(salary)

	for k, v := range salary {
		fmt.Printf("Key: %s and the value is: %d\n", k, v)
	}

	for _, v := range salary {
		fmt.Printf("The salary is: %d\n", v)
	}

	// sal := make(map[string]int)
	// sal1 := map[string]int{}

}