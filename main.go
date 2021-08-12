package main

import (
	"fmt"
)

func main() {

	// fmt.Print(v)

	var getMinMax = func(n []float32) (float32, float32) {
		var min, max float32
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	m := make(map[string]float32)

	m["Henry"] = 90
	m["Dyah"] = 85
	m["Puti"] = 78.5
	m["Reza"] = 83
	m["Rahma"] = 69

	v := make([]float32, 0, len(m))

	for _, value := range m {
		v = append(v, value)
	}

	var min, max = getMinMax(v)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", m, min, max)
}
