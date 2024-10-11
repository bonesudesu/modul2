package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Print("Masukkan nilai k: ")
	fmt.Scanln(&k)

	var hasil float64 = 1
	for i := 0; i < k; i++ {
		hasil *= (4*float64(i) + 2) * (4*float64(i) + 2) / ((4*float64(i) + 1) * (4*float64(i) + 3))
	}

	fmt.Printf("Hampiran akar 2 untuk k = %d adalah %.10f\n", k, hasil)
}
