package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	fmt.Println("Bilangan:", bilangan)
	fmt.Print("Faktor: ")

	var faktor int
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
			faktor++
		}
	}

	if faktor == 2 {
		fmt.Println("\nPrima: true")
	} else {
		fmt.Println("\nPrima: false")
	}
}
