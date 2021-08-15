package main

import (
	"fmt"
	"tugas-dua/sorting"
)

func main() {

	// Deklarasi variabel Map
	nilaiSiswa := make(map[string]float32)

	// Inisialisasi variabel Map
	nilaiSiswa["Henry"] = 90
	nilaiSiswa["Dyah"] = 85
	nilaiSiswa["Puti"] = 78.5
	nilaiSiswa["Reza"] = 83
	nilaiSiswa["Rahma"] = 69

	// Sorting Nilai Rata - Rata
	var nilaiRata float32 = sorting.GetMapAverage(nilaiSiswa)
	fmt.Println(nilaiRata)

	// Sorting Nilai Minimal
	var nilaiMinimal float32 = sorting.GetMapMinValue(nilaiSiswa)
	fmt.Println(nilaiMinimal)

	// Sorting Nilai Maximal
	var nilaiMaximal float32 = sorting.GetMapMaxValue(nilaiSiswa)
	fmt.Println(nilaiMaximal)

}
