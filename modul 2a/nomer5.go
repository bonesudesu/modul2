package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan sebuah kalimat: ")
	scanner.Scan()
	kalimat := scanner.Text()

	fmt.Println("Kalimat yang dimasukkan adalah: ", kalimat)

	// menghitung jumlah karakter dalam kalimat
	jumlahKarakter := len(kalimat)
	fmt.Println("Jumlah karakter dalam kalimat: ", jumlahKarakter)

	// menghitung jumlah kata dalam kalimat
	jumlahKata := 0
	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] == ' ' {
			jumlahKata++
		}
	}
	jumlahKata++ // menambahkan jumlah kata terakhir
	fmt.Println("Jumlah kata dalam kalimat: ", jumlahKata)

	// menghitung jumlah huruf kecil dalam kalimat
	jumlahHurufKecil := 0
	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] >= 'a' && kalimat[i] <= 'z' {
			jumlahHurufKecil++
		}
	}
	fmt.Println("Jumlah huruf kecil dalam kalimat: ", jumlahHurufKecil)

	// menghitung jumlah huruf besar dalam kalimat
	jumlahHurufBesar := 0
	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] >= 'A' && kalimat[i] <= 'Z' {
			jumlahHurufBesar++
		}
	}
	fmt.Println("Jumlah huruf besar dalam kalimat: ", jumlahHurufBesar)
}
