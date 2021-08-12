package main

import "fmt"

func main() {
	nilai := make(map[string]float32)

	nilai["Henry"] = 90
	nilai["Dyah"] = 85
	nilai["Puti"] = 78.5
	nilai["Reza"] = 83
	nilai["Rahma"] = 69
	fmt.Println("Nama dan Nilai Siswa:", nilai)

	// Nilai Terendah dan Tertinggi
	var min float32 = -1
	var max float32
	for _, e := range nilai {
		if min == -1 || (e >= 0 && min > e) {
			min = e
		}
		if max < e {
			max = e
		}
	}
	fmt.Println("Niali Terendah:", min)
	fmt.Println("Nilai Tertinggi:", max)

	//Nilai rata-rata
	var sum float32
	for _, v := range nilai {
		sum += v
	}
	rataRata := float32(sum) / float32(len(nilai))
	fmt.Println("Nilai rata-rata:", rataRata)
}
