package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("siswa.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened siswa.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our data array
	siswa := make(map[string]float64)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into siswa which we defined above
	json.Unmarshal([]byte(byteValue), &siswa)
	fmt.Println(siswa)

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
