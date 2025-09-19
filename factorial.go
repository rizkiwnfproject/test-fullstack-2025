package main

import (
	"fmt"
	"math"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func f(n int) int {
	result := float64(factorial(n)) / float64(int(math.Pow(2, float64(n))))
	return int(math.Ceil(result))
}

func main() {
	var n int
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&n)

	fmt.Printf("f(%d) = %d\n", n, f(n))
}
