package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	frequency := make(map[int]int)

	// Menghitung frekuensi kemunculan setiap angka
	for _, char := range angka {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			frequency[num]++
		}
	}

	// Memilih angka-angka yang hanya muncul sekali
	var uniqueNumbers []int
	for _, char := range angka {
		num, _ := strconv.Atoi(string(char))
		if frequency[num] == 1 {
			uniqueNumbers = append(uniqueNumbers, num)
		}
	}

	return uniqueNumbers
}

func main() {
	fmt.Println(munculSekali("9870987"))
	fmt.Println(munculSekali("1423714239"))
	fmt.Println(munculSekali("98765"))
	fmt.Println(munculSekali("559977881122"))
	fmt.Println(munculSekali("00776629930045"))
}
