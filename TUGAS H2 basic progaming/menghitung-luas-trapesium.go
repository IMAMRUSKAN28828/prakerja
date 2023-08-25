package main

import "fmt"

func main() {
    var tinggi, panjangSisiAtas, panjangSisiBawah float64

    fmt.Print("Masukkan tinggi trapesium: ")
    fmt.Scanln(&tinggi)

    fmt.Print("Masukkan panjang sisi atas trapesium: ")
    fmt.Scanln(&panjangSisiAtas)

    fmt.Print("Masukkan panjang sisi bawah trapesium: ")
    fmt.Scanln(&panjangSisiBawah)

    luas := 0.5 * (panjangSisiAtas + panjangSisiBawah) * tinggi
    fmt.Printf("Luas trapesium adalah: %f\n", luas)
}
