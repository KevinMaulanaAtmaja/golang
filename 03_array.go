package main

import "fmt"

func array() {
	fmt.Println("=============ARRAY=============")
	// kumpulan data dg ukuran ttap dan tipe data yg sama
	// exp: 12 telur g bisa ditmbh/dikurangi
	// index=> no urut utk akses elemen di array, slalu dimulai dari 0
	var angka [3]int = [3]int{10, 20, 30}
	fmt.Println(angka)    //[10 20 30]
	fmt.Println(angka[1]) //20
	angka[2] = 60
	fmt.Println("Angka yg telah diubah:", angka[2])

	fmt.Println("=============FUNCTION ARRAY=============")
	number := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Jumlah elemen:", len(number))       // [10 20 30 40 50]
	fmt.Println("Index ke-2:", number[2])            //30 ->akses elemen
	number[2] = 65                                   //ubah elemen di index ke-2
	fmt.Println("Nilai yg telah diubah:", number[2]) //65

	fmt.Println("Ini adalah aaray:")
	for index, value := range number {
		fmt.Println("Isi index ke:", index, "=", value)
	}

	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{4, 5, 6}

	fmt.Println("arr1 == arr2:", arr1 == arr2) //false
	fmt.Println("arr1 != arr2:", arr1 != arr2) //true
}
