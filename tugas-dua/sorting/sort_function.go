package sorting

import (
	"fmt"
)

// Fungsi Mendapatkan Nilai Rata-Rata Pada Data Map
func GetMapAverage(values map[string]float32) float32 {
	var result float32
	for _, value := range values {
		result += value
	}

	fmt.Printf("Nilai Rata - Rata : ")
	return result / float32(len(values))
}

// Fungsi Mendapatkan Nilai Terendah Pada Data Map
func GetMapMinValue(values map[string]float32) float32 {
	var minValue float32
	var name string
	for key, value := range values {
		if minValue == 0 || value < minValue {
			minValue = value
			name = key
		}
	}
	fmt.Printf("Nilai Terendah %s : ", name)
	return minValue
}

// Fungsi Mendapatkan Nilai Tertinggi Pada Data Map
func GetMapMaxValue(values map[string]float32) float32 {
	var maxValue float32
	var name string
	for key, value := range values {
		if maxValue == 0 || value > maxValue {
			maxValue = value
			name = key
		}
	}
	fmt.Printf("Nilai Tertinggi %s : ", name)
	return maxValue
}
