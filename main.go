package main

import "fmt"

/*
	Soal :
	1. Henry  : 90
	2. Dyah : 85
	3. Puti : 78.5
	4. Reza : 83
	5. Rahma : 69

	buat maps, dengan declare make, mencari
	1. Nilai tertinggi
	2. Nilai rata-rata
	3. Nilai terendah
*/

func main() {

	nilai := make(map[string]float64)
	nilai["Henry"] = 90
	nilai["Dyah"] = 85
	nilai["Puti"] = 78.5
	nilai["Reza"] = 83
	nilai["Rahma"] = 69

	var max, min, sum float64

	//Nilai Tertinggi
	for _, v := range nilai {
		if v > max {
			max = v
		}
	}

	fmt.Println("Nilai tertinggi \t:", max)

	//Rata-rata
	for _, v := range nilai {
		sum += v
	}
	avg := sum / float64(len(nilai))
	fmt.Println("Nilai rata-ratanya \t:", avg)

	//Nilai Terendah
	min = nilai["Henry"]
	for _, v := range nilai {
		if v < min {
			min = v
		}
	}
	fmt.Println("Nilai terendah \t\t:", min)

}
