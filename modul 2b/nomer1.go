package main

import "fmt"

func main() {
	var warnaGelas1, warnaGelas2, warnaGelas3, warnaGelas4 string
	var hasilPercobaan []string

	for i := 0; i < 5; i++ {
		fmt.Println("Percobaan ke-", i+1)
		fmt.Print("Warna Gelas 1: ")
		fmt.Scanln(&warnaGelas1)
		fmt.Print("Warna Gelas 2: ")
		fmt.Scanln(&warnaGelas2)
		fmt.Print("Warna Gelas 3: ")
		fmt.Scanln(&warnaGelas3)
		fmt.Print("Warna Gelas 4: ")
		fmt.Scanln(&warnaGelas4)

		hasilPercobaan = append(hasilPercobaan, fmt.Sprintf("%s,%s,%s,%s", warnaGelas1, warnaGelas2, warnaGelas3, warnaGelas4))
	}

	for _, percobaan := range hasilPercobaan {
		warna := []string{}
		for _, s := range percobaan {
			if s == ',' {
				continue
			}
			warna = append(warna, string(s))
		}

		if warna[0] == "merah" && warna[1] == "kuning" && warna[2] == "hijau" && warna[3] == "ungu" {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}
