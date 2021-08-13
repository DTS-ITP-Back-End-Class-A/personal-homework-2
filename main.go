package main

import "fmt"

func main() {
	siswa := make(map[string]float64)
	siswa["Henry"] = 90
	siswa["Dyah"] = 85
	siswa["Puti"] = 78.5
	siswa["Reza"] = 83
	siswa["rahma"] = 69

	getMaxNumber(siswa)
	averageNumber(siswa)
	getMinNumber(siswa)
}

func getMaxNumber(student map[string]float64) {
	var maxNum float64
	var namaSiswa string
	for key, value := range student {
		if value > maxNum {
			maxNum = value
			namaSiswa = key
		}
	}
	fmt.Println("Nilai Tertinggi Adalah ", maxNum, " yang diperoleh Oleh ", namaSiswa)
}

func averageNumber(student map[string]float64) {
	var sum float64
	for _, value := range student {
		sum += value
	}
	ave := sum / float64(len(student))
	fmt.Println("Nilai rata-rata siswa adalah", ave)
}

func getMinNumber(student map[string]float64) {
	var minNum float64
	var namaSiswa string
	for key, value := range student {
		if key == "Henry" {
			minNum = value
		}
		if value < minNum {
			minNum = value
			namaSiswa = key
		}

	}
	fmt.Println("Nilai Terendah Adalah ", minNum, " yang diperoleh Oleh ", namaSiswa)
}
