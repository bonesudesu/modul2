package main

import "fmt"

func main() {
	// Contoh data parsel
	parcels := []int{8500, 9250, 11750}

	// Iterasi melalui data parsel
	for i, parcelWeight := range parcels {
		// Hitung berat parsel dalam kg dan gram
		kg := parcelWeight / 1000
		gr := parcelWeight % 1000

		// Hitung biaya kirim
		baseCost := kg * 10000 // Biaya dasar per kg

		var extraCost int // Biaya tambahan untuk sisa berat
		if gr >= 500 {
			extraCost = gr * 5 // Biaya tambahan per gram
		} else {
			extraCost = gr * 15 // Biaya tambahan per gram
		}

		// Bebaskan biaya sisa berat jika total berat > 10kg
		if kg > 10 {
			extraCost = 0
		}

		// Hitung total biaya kirim
		totalCost := baseCost + extraCost

		// Tampilkan hasil
		fmt.Printf("%d Contoh #%d\n", i+1, i+1)
		fmt.Printf("Berat parsel (gram): %d\n", parcelWeight)
		fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", baseCost, extraCost)
		fmt.Printf("Total biaya: Rp. %d\n\n", totalCost)
	}
}
