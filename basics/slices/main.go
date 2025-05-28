package main

import (
	"fmt"
)

func main() {
	// a[low : high]
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	// fmt.Println(primes)

	// for i := 0 ; i < len(primes); i++ {
	// 	fmt.Println(primes[i])
	// }

	// for i, v := range primes {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, v)
	// }

	// var i int = 0
	// // Using a while-like loop
	// for i < len(primes) {
	// 	fmt.Println("Index:", i, "Value:", primes[i])
	// 	i++
	// }

	// Slicing
	slice1 := primes[2:5] // 5 is exclusive
	fmt.Println("Slice1:", slice1)

	slice2 := primes[5:] // From index 5 to the end
	fmt.Println("Slice2:", slice2)

	slice3 := primes[:5]
	fmt.Println("Slice3:", slice3)

	slice4 := primes[:] // Full slice
	fmt.Println("Slice4:", slice4)

}

