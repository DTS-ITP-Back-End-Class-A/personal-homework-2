package main

import (
	"fmt"
)

func main() {
	// inisialisai
	var students = make(map[string]float32)

	students["Henry"] = 90
	students["Dyah"] = 85
	students["Puti"] = 78.5
	students["Reza"] = 83
	students["Rahma"] = 69

	var minValue float32 = 0
	var maxValue float32 = 0
	var average float32 = getAverage(students)

	for _, value := range students {
		if minValue == 0 || value < minValue {
			minValue = value
		}

		if value > maxValue {
			maxValue = value
		}
	}

	fmt.Println(" Nilai tertinggi", maxValue)
	fmt.Println(" Nilai terendah", minValue)
	fmt.Println(" Nilai rata-rata", average)

}

func getAverage(values map[string]float32) float32 {
	var result float32 = 0

	for _, value := range values {
		result += value
	}

	return result / float32(len(values))
}
