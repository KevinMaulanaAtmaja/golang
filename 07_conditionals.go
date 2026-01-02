package main

import "fmt"

func conditionals() {
	fmt.Println("=============CONDITIONALS=============")
	// salah satu cara utk buat program ambil keputusan brdsrkan logika (benar/salah)
	// jika benar blok kode if jalan, jika false blok kode tersbt masuk ke else

	// format
	// if kondisi1 {
	// 	// jika kondisi1 true
	// } else if kondisi2 {
	// 	// jika kondisi2 true
	// } else {
	// 	// jika smua false
	// }

	// contoh
	// if
	fmt.Println("=============IF=============")
	x := 10
	if x > 5 {
		fmt.Println("x lebih besar dari 5")
	}

	fmt.Println("=============IF-ELSE=============")
	// if-else
	y := 3
	if y > 5 {
		fmt.Println("x lebih besar dari 5")
	} else {
		fmt.Println("x tidak lebih besar dari 5")
	}

	fmt.Println("=============IF-ELSE_IF-ELSE=============")
	// if-else_if-else
	z := 11
	if z > 10 {
		fmt.Println("x lebih besar dari 10")
	} else if z > 5 {
		fmt.Println("x lebih besar dari 5")
	} else {
		fmt.Println("x tidak lebih besar dari 5")
	}

	fmt.Println("==========================")
	//
	var nilai int
	fmt.Print("Masukkan nilai ujian:")
	fmt.Scanln(&nilai)

	if nilai >= 90 {
		fmt.Println("Nilai Anda A (Sangat Baik).")
	} else if nilai >= 75 {
		fmt.Println("Nilai Anda B (Baik).")
	} else if nilai >= 60 {
		fmt.Println("Nilai Anda C (Cukup).")
	} else {
		fmt.Println("Nilai Anda D (Perlu Perbaikan).")
	}
}
