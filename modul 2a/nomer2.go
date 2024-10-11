package main

import "fmt"

func tahunKabisat(tahun int) bool {
	if tahun%4 == 0 {
		if tahun%100 == 0 {
			if tahun%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	if tahunKabisat(tahun) {
		fmt.Printf("%d adalah tahun kabisat\n", tahun)
	} else {
		fmt.Printf("%d bukan tahun kabisat\n", tahun)
	}
}
