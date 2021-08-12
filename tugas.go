package main

import "fmt"

func main() {

	/*
		1. Henry : 90
		2. Dyah : 85
		3. Puti : 78.5
		4. Reza : 83
		5. Rahma : 69

		buat maps, dengan declare make, mencari
		- Nilai tertinggi
		- Nilai rata2
		- Nilai terendah
	*/

	// buat maps, dengan declare make
	siswa := make(map[string]float64)

	// assign "key" dan valuenya
	siswa["Henry"] = 90
	siswa["Dyah"] = 85
	siswa["Puti"] = 78.5
	siswa["Reza"] = 83
	siswa["Rahma"] = 69

	// coba cek dulu mapnya
	fmt.Println(siswa)

	// 1. cari nilai tertinggi
	// tampung dulu nilai keynya sbg acuan di variable
	max := siswa["Henry"]

	// kita iterasi mapnya dengan for utk mencari valuenya
	for _, v := range siswa {

		// setelah dapet valuenya, buat logikanya
		if v > max {
			max = v
		}
	}

	// coba print bener gak?
	fmt.Println("ini nilai tertinggi", max)

	// 2. cari nilai rata2
	// cari dulu jumlah semuanya
	// dg cara buat dulu var utk menampung jumlah seluruh value
	var sum float64 = 0

	// kita iterasi mapnya dengan for utk mencari valuenya
	for _, v := range siswa {

		// valuenya kita jumlahkan dg sum
		sum += v
	}

	// coba print bener gak?
	fmt.Println("ini nilai sum", sum)

	// tinggal cari rata2nya dg dibagi jumlah seluruh siswa
	var avg float64 = sum / float64(len(siswa))

	// coba print bener gak?
	fmt.Println("ini nilai avg", avg)

	// 3. cari nilai terendah
	// tampung dulu nilai keynya sbg acuan di variable
	min := siswa["Henry"]

	// kita iterasi mapnya dengan for utk mencari valuenya
	for _, v := range siswa {

		// setelah dapet valuenya, buat logikanya
		if v < min {
			min = v
		}
	}

	// coba print bener gak?
	fmt.Println("ini nilai terendah", min)

}
