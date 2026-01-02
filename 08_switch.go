package main

import "fmt"

func switchs() {
	fmt.Println("=============SWITCH=============")
	// cocok utk cek nilai dari satu variable dg bnyk kasus
	// otomatis break, pake fallthrough utk lanjut ke case lain
	// default dijalnkan jika tdk adayg cocok

	// numeric
	angka := 1
	switch angka {
	case 1:
		fmt.Println("Angka satu")
	case 2:
		fmt.Println("Angka dua")
	case 3:
		fmt.Println("Angka tigas")
	default:
		fmt.Println("Angka tidak dikenal")
	}

	// string (case-sensitive)
	hari := "Sabtu"
	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari Kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari Libur")
	default:
		fmt.Println("Hari tidak valid")
	}

	// short statement=> utk deklarasi var smentara, spya lbih ringkas
	// switch v := expression; v {
	// case value1:
	// aksi
	// case value2:
	// aksi
	// default:
	// aksi default
	// }
	switch panjang := len("Geologi"); {
	case panjang >= 10:
		fmt.Println("Kata terlalu panjang")
	case panjang >= 5:
		fmt.Println("Kata sedang")
	default:
		fmt.Println("Kata pendek")
	}
}
