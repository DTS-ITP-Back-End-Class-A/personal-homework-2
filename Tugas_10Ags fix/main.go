package main

import (
	"fmt"
)

/* Tugas_10Ags

Henry: 90
Dyah: 85
Puti: 78.5
Reza: 83
Rahma: 69

- maps dengan make
- nilai tertinggi
- nilai rata-rata
- nilai terendah

*/

func main(){

	//list maps dengan make.
	daftar := make(map[string]float64)

	daftar["Henry"] = 90
	daftar["Dyah"] = 85
	daftar["Puti"] = 78.5
	daftar["Reza"] = 83
	daftar["Rahma"] = 69

	//cetak list.
	fmt.Println("Daftar Nilai:", daftar)

	//Proses perulangan dan percabangan untuk mencari nilai tertinggi dan terendah
	var max float64 = daftar["Puti"]
    min := daftar["Puti"] //ini adalah data untuk perbandingan, bisa pakai data yang mana saja.
    for _, nilai := range daftar {
        if max < nilai {
            max = nilai //jadi nilai ini akan berubah sebanyak data yang ada, data tertinggi yang bertahan.
        }
        if min > nilai {
            min = nilai //casenya sama dengan max namun data yang bertahan adalah data terendah.
        }
    }

	//cetak nilai tertinggi dan terendah.
	fmt.Println("Nilai tertinggi: ", max)
	fmt.Println("Nilai terendah: ", min)

	//Proses perulangan untuk menambahkan isi data keseluruhan, agar bisa dibagi dengan banyaknya data.
	var total float64 = 0
	var average float64 
	for _, data := range daftar {  
		total = total + data  
	}  
	average = total / float64(len(daftar)) 

	//cetak hasil rata-rata.
	fmt.Println("Nilai Rata-rata: ", average)

}