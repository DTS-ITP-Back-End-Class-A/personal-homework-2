package main

import (
	"fmt"
)

func main() {

	siswa := make(map[string]float64)
	siswa["Henry"] = 90
	siswa["Henry"] = 90
	siswa["Dyah"] = 85
	siswa["Puti"] = 78.5
	siswa["Reza"] = 83
	siswa["Rahma"] = 69

	//fmt.Println(siswa)

	var temp_nilai, temp_rata, rata, max, min float64
	temp_nilai = 0
	temp_rata = 0
	rata = 0
	max = 0
	min = 0
	for _, val := range siswa {
		if val > temp_nilai {
			temp_nilai = val
			max = temp_nilai
		}
		temp_rata += val
		//fmt.Println(val)
	}

	for _, val := range siswa {
		if val < temp_nilai {
			temp_nilai = val
			min = temp_nilai
		}
	}

	rata = temp_rata / float64(len(siswa))
	fmt.Println("Nilai Tertinggi = ", max)
	fmt.Println("Nilai rata-rata = ", rata)
	fmt.Println("Nilai Terendah = ", min)
}
