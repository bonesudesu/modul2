package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var bunga []string
	for {
		fmt.Print("Masukkan nama bunga (ketik 'SELESAI' untuk berhenti): ")
		var nama string
		fmt.Fscanln(reader, &nama)

		if strings.ToUpper(nama) == "SELESAI" {
			break
		}

		bunga = append(bunga, nama)
	}

	pita := strings.Join(bunga, " - ")
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Jumlah bunga: %d\n", len(bunga))
}
