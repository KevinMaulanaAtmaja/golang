package defpanrec

import "fmt"

func Panics() {
	fmt.Println("=============PANIC=============")
	// panic=> utk hentikan eksekusi normal dg error fatal, jalankan smua defer(stack unwinding) lalu program berhenti, kecuali ada recover
	// knp dipake? kondisi luar biasa yg tdk bisa ditangani scara normal (corrupt state)
	// panic("error message")
	defer fmt.Println("defer ini ttap jalan sblum program mati")
	fmt.Println("Sebelum panic")
	panic("ada sesuatu yang fatal")
	fmt.Println("Setelah panic") //baris ini tdk dieksekusi
}
