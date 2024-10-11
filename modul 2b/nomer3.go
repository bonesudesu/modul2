package main

import (
	"fmt"
)

func main() {
	var (
		beratKiri, beratKanan float64
	)

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKiri, &beratKanan)

		if beratKiri <= 9 || beratKanan <= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		}

		if beratKiri+beratKanan > 150 || beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
