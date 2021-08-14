package main

import "fmt"

func main() {
	siswa := make(map[string]float64)
	siswa["Henry"] = 90
	siswa["Dyah"] = 85
	siswa["Puti"] = 78.5
	siswa["Reza"] = 83
	siswa["rahma"] = 69

	getMaxMinAveNumber(siswa)

}

func getMaxMinAveNumber(student map[string]float64) {
	var maxNum float64
	var minNum float64
	var sum float64
	var namaSiswaMax string
	var namaSiswaMin string
	for key, value := range student {
		if value > maxNum {
			maxNum = value
			namaSiswaMax = key
		} else {
			minNum = value
			namaSiswaMin = key
		}

		sum += value

	}
	ave := sum / float64(len(student))

	fmt.Println("Nilai Tertinggi Adalah ", maxNum, " yang diperoleh Oleh ", namaSiswaMax)
	fmt.Println("Nilai rata-rata siswa adalah", ave)
	fmt.Println("Nilai Terendah Adalah ", minNum, " yang diperoleh Oleh ", namaSiswaMin)
}

// func averageNumber(student map[string]float64) {
// 	var sum float64
// 	for _, value := range student {
// 		sum += value
// 	}
// 	ave := sum / float64(len(student))
// 	fmt.Println("Nilai rata-rata siswa adalah", ave)
// }

// func getMinNumber(student map[string]float64) {
// 	var minNum float64
// 	var namaSiswa string
// 	for key, value := range student {
// 		if key == "Henry" {
// 			minNum = value
// 		}
// 		if value < minNum {
// 			minNum = value
// 			namaSiswa = key
// 		}

// 	}
// 	fmt.Println("Nilai Terendah Adalah ", minNum, " yang diperoleh Oleh ", namaSiswa)
// }
