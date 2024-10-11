package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nam = "A"
	} else if nam > 72.5 {
		nam = "AB"
	} else if nam > 65 {
		nam = "B"
	} else if nam > 57.5 {
		nam = "BC"
	} else if nam > 50 {
		nam = "C"
	} else if nam > 40 {
		nam = "D"
	} else { // Simplified 'else if' for clarity
		nam = "E"
	}
	nmk = nam // Assign 'nam' to 'nmk'
	fmt.Println("Nilai mata kuliah: ", nmk)
}
