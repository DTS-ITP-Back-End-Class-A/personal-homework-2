package main

import "fmt"

func main() {
	// daftar nama dan nilai mahasiswa
	mahasiswa := make(map[string]float32)
	mahasiswa["henry"] = 90
	mahasiswa["dyah"] = 85
	mahasiswa["puti"] = 78.5
	mahasiswa["reza"] = 83
	mahasiswa["rahma"] = 69

	// map slices mahasiswa

	fmt.Println(mahasiswa)

	var jumlah, maksi, minim float32

	// Nilai Rata-Rata
	for _, nilai := range mahasiswa {
		jumlah += nilai
	}
	rata := float32(jumlah) / float32(len(mahasiswa))
	fmt.Println("Nilai Rata-Rata adalah:", rata)

	// Nilai Terbesar
	for _, a := range mahasiswa {
		if a > maksi {
			maksi = a
		}
	}
	fmt.Println("Nilai Terbesar adalah :", maksi)

	// Nilai Terkecil
	for _, b := range mahasiswa {
		if b < minim {
			minim = b
		}
	}
	fmt.Println("Nilai Terkecil adalah :", minim)

}
