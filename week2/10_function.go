package week2

import "fmt"

func Functions() {
	fmt.Println("=============FUNCTION=============")
	// blok kode utk jalankan tugas trtentu
	// utk hindari duplikasi kode, lbih terstruktur, mudah dibaca dan dipelihara
	// didefinisikan pake keyword "func"
	// param-> var yg dikirimkan ke func, supaya bisa pake data yg beda2
	// variadidc func-> mnrima jmlh argumen yg unlimited, pake tanda ...

	fmt.Println("=============FUNCTION WITHOUT PARAM=============")
	// pemanggilan fungsi
	greeting()
	greeting()
	fmt.Println("=============FUNCTION WITH PARAM=============")
	// pake satu param
	greet("Kevin")
	greet("Amel")
	// pake multiple param
	greetAdd("Malam", "Dina")
	fmt.Println("=============FUNCTION WITH RETURN=============")
	// pake satu return
	hasil := tambah(10, 20)
	fmt.Println(hasil)
	fmt.Println(tambah(30, 40))
	
	// pake multiple return
	hasilBagi, pesan := bagi(10, 2)
	fmt.Println("Hasil:", hasilBagi, "| Status:", pesan)
	fmt.Println(bagi(10, 0))
	
	// pake named return values
	fmt.Println(kali(3, 3))
	
	fmt.Println("=============VARIADIC FUNCTION=============")
	// pake variadic func
	fmt.Println(jumlahAngka(1, 9, 5))
	fmt.Println(jumlahAngka(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(jumlahAngka())
}

// tanpa param
func greeting() {
	fmt.Println("Halo, selamat datang!")
}

// pake param
func greet(name string) {
	fmt.Println("Konnichiwa,", name, "Arigatogozaimasu.")
}

// pake multiple param
func greetAdd(waktu, nama string) {
	fmt.Println("Selamat", waktu, ", "+nama)
}

// pake return
func tambah(a int, b int) int {
	return a + b
}

// pake multiple return
func bagi(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Tidak bisa bagi dg nol"
	}
	return a / b, "Berhasil"
}

// memberi nama pada return (named return values)
func kali(a, b float64) (hasil float64) {
	hasil = a * b
	// ckp tulis return aja
	return
}

// variadic func
func jumlahAngka(angka ...int) int {
	total := 0
	for _, value := range angka {
		total += value //total = total + value
	}
	return total
}
