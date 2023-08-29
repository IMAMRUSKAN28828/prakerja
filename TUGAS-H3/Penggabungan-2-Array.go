package main

import "fmt"

func arrayMerge(nama [5]string, asal [5]string, ubahNama string, asalBaru string) [5]string {
	for i := 0; i < len(nama); i++ {
		if nama[i] == ubahNama {
			asal[i] = asalBaru
			break
		}
	}
	return asal
}

func main() {
	nama := [5]string{
		"Robi", "Romi", "Roni", "Ridwan", "Rafli",
	}

	asal := [5]string{
		"Surabaya", "Jakarta", "Bandung", "Jombang", "Malang",
	}

	ubahNama := "" // Var untuk merubah nama 
	asalBaru := "" // Var untuk merubah asal

	asal = arrayMerge(nama, asal, ubahNama, asalBaru)

	for i := 0; i < len(nama); i++ {
		fmt.Printf("Nama: %s, Asal: %s\n", nama[i], asal[i])
	}
}