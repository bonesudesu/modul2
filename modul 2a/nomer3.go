package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari int
	fmt.Print("Jejari = ")
	fmt.Scanln(&jariJari)

	volume := (4.0 / 3.0) * math.Pi * float64(jariJari) * float64(jariJari) * float64(jariJari)
	luas := 4 * math.Pi * float64(jariJari) * float64(jariJari)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", jariJari, volume, luas)
}
