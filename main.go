package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	jsonFile, err := os.Open("students.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	studentScore := jsonFileToMap(jsonFile)

	max, min, avg := calculateScore(studentScore)
	fmt.Println(max, min, avg)
}

func jsonFileToMap(jsonFile *os.File) map[string]float64 {
	byteValue, _ := ioutil.ReadAll(jsonFile)
	data := make(map[string]float64)

	json.Unmarshal([]byte(byteValue), &data)

	return data
}

func calculateScore(data map[string]float64) (maxScore, minScore, avgScore string) {
	max := float64(0) // default score
	maxKey := ""
	min := data["Henry"] // random pick
	minKey := ""
	sum := float64(0)

	for k, v := range data {
		if v > max {
			max = v
			maxKey = k
		}

		if v < min {
			min = v
			minKey = k
		}

		sum += v
	}

	maxScore = fmt.Sprintf("Nilai tertinggi diperoleh %s dengan score %.0f", maxKey, max)
	minScore = fmt.Sprintf("\nNilai terendah diperoleh %s dengan score %.0f", minKey, min)
	avgScore = fmt.Sprintf("\nNilai rata-rata kelas adalah %.1f", sum/float64(len(data)))

	return
}
