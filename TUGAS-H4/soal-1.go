package main

import "fmt"

type TipeMobil int

const (
    Sedan TipeMobil = iota
    SUV
    Truk
    // Tambahkan tipe mobil lainnya jika diperlukan
)

type Mobil struct {
    Tipe       TipeMobil
    BahanBakar float64
}

func (m Mobil) HitungJarakPerkiraan() float64 {
    EfisiensiBahanBakar := 1.5  // Efisiensi bahan bakar dalam kilometer per liter
    JarakPerkiraan := EfisiensiBahanBakar * m.BahanBakar
    return JarakPerkiraan
}

func main() {
    inovaReborn := Mobil{
        Tipe:       Sedan,
        BahanBakar: 40.0,
    }
    
    inovaReborn.BahanBakar = 120.0 // Mengubah bahan bakar menjadi 120.0
    jarakPerkiraan := inovaReborn.HitungJarakPerkiraan()
    fmt.Printf("Perkiraan jarak yang bisa ditempuh dengan Inova Reborn: %.2f km\n", jarakPerkiraan)
}
