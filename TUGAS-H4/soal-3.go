package main

import (
	"fmt"
)

func cariMaksimumDanMinimum(angka []int) (*int, *int) {
	maksimum := angka[0]
	minimum := angka[0]

	for _, num := range angka {
		if num > maksimum {
			maksimum = num
		}
		if num < minimum {
			minimum = num
		}
	}

	return &maksimum, &minimum
}

func main() {
	var angkaInput [6]int

	fmt.Println("Masukkan 6 angka:" )
	for i := 0; i < 6; i++ {
		fmt.Scan(&angkaInput[i])
	}

	// Mendapatkan alamat dari nilai-nilai maksimum dan minimum
	alamatMaksimum, alamatMinimum := cariMaksimumDanMinimum(angkaInput[:])

	// Dereferensi alamat untuk mendapatkan nilai sebenarnya
	maksimum := *alamatMaksimum
	minimum := *alamatMinimum

	fmt.Println("Nilai Maksimum:", maksimum)
	fmt.Println("Nilai Minimum:", minimum)
}
