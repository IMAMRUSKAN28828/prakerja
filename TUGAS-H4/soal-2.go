package main

import (
	"fmt"
	"math"
)

type Siswa struct {
	Nama  string
	Skor  int
}

type Sekolah struct {
	Siswas []Siswa
}

func (s *Sekolah) TambahSiswa(nama string, skor int) {
	siswa := Siswa{Nama: nama, Skor: skor}
	s.Siswas = append(s.Siswas, siswa)
}

func (s Sekolah) RataRataSkor() float64 {
	totalSkor := 0
	for _, siswa := range s.Siswas {
		totalSkor += siswa.Skor
	}
	return float64(totalSkor) / float64(len(s.Siswas))
}

func (s Sekolah) SkorTerendahTertinggi() (int, int) {
	skorTerendah := math.MaxInt32
	skorTertinggi := math.MinInt32

	for _, siswa := range s.Siswas {
		if siswa.Skor < skorTerendah {
			skorTerendah = siswa.Skor
		}
		if siswa.Skor > skorTertinggi {
			skorTertinggi = siswa.Skor
		}
	}

	return skorTerendah, skorTertinggi
}

func main() {
	sekolah := Sekolah{}

	// Mengisi data siswa
	sekolah.TambahSiswa("Robi", 85)
	sekolah.TambahSiswa("Roni", 70)
	sekolah.TambahSiswa("Romi", 90)
	sekolah.TambahSiswa("Ridwan", 75)
	sekolah.TambahSiswa("Ronaldo", 50)

	// Menampilkan informasi
	fmt.Printf("Rata-rata skor: %.2f\n", sekolah.RataRataSkor())

	skorTerendah, skorTertinggi := sekolah.SkorTerendahTertinggi()
	fmt.Printf("Skor terendah: %d\n", skorTerendah)
	fmt.Printf("Skor tertinggi: %d\n", skorTertinggi)
}
