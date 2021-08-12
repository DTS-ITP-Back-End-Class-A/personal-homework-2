package main

import (
	"fmt"
)

func main() {
	// Membuat map nilai_siswa dengan make
	nilai_siswa := make(map[string]float64)

	nilai_siswa["Henry"] = 90
	nilai_siswa["Dyah"] = 85
	nilai_siswa["Puti"] = 78.5
	nilai_siswa["Reza"] = 83
	nilai_siswa["Rahma"] = 69

	//Mencari nilai tertinggi
	var nilai_tertinggi float64
	for _, v := range nilai_siswa {
		if v > nilai_tertinggi {
			nilai_tertinggi = v
		}
	}
	fmt.Println("Nilai Tertinggi : ", nilai_tertinggi)

	//Mencari nilai terendah
	var nilai_terendah float64 = nilai_tertinggi
	for _, v := range nilai_siswa {
		if v < nilai_terendah {
			nilai_terendah = v
		}
	}
	fmt.Println("Nilai Terendah : ", nilai_terendah)

	//Mencari nilai rata - rata
	var jumlah_nilai float64
	for _, v := range nilai_siswa {
		jumlah_nilai = jumlah_nilai + v
	}
	fmt.Println("Nilai rata - rata siswa : ", jumlah_nilai/float64((len(nilai_siswa))))

}
