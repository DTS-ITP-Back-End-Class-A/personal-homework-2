/*
	1. Henry : 90
	2. Dyah : 85
	3. Puti : 78.5
	4. Reza : 83
	5. Rahma : 69

	buat maps, dengan declare make, mencari
	- Nilai Tertinggi
	- Nilai rata2
	- Nilai terendah
*/
package main

import (
	"fmt"
	// "math"
)

func main() {
	var sum float64

	students := make(map[string]float64)
	students["Henry"] = 90
	students["Dyah"] = 85
	students["Puti"] = 78.5
	students["Reza"] = 83
	students["Rahma"] = 69
	fmt.Println("daftar nama dan nilai siswa :", students)

	//rata rata untuk siswa
	for _, v := range students {
		sum += v

	}
	avg := (float64(sum) / float64(len(students)))
	fmt.Println("rata rata siswa : ", avg)

	//nilai tertinggi
	var largeNumber, minimumNumber, temp float64
	for _, value := range students {
		if value > temp {
			temp = value
			largeNumber = temp

		} 
	}

	fmt.Println("maximum number", largeNumber)

	for _, value := range students {
		if value < temp {
			temp = value
			minimumNumber = temp

		}
	}

	fmt.Println("minimum number", minimumNumber)

}


