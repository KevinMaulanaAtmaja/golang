package defpanrec

import "fmt"

func Recoverys() {
	fmt.Println("=============RECOVERY=============")
	// recover=> utk menangkap panic hanya bila dipanggil didlm fungsi yg didefer, jika berhasil ditangkap, eksekusi program bisa lanjut
	// aturan kunci: hnya efektid didlm defer, r := recover() akan bernilai nil jika tdk ada panic
	fmt.Println("Program mulai")
	bagi(10, 2) //normal
	bagi(5, 0) //ini akan terjadi panic, tpi recover akan menangkapnya
	fmt.Println("Program selesai dengan aman")
}

func tanganiPanic() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi panic, tapi sudah ditangani:", r)
	}
}

func bagi(a, b int) {
	defer tanganiPanic() //defer akan pastikan tanganiPanic dipanggil walaupun panic
	fmt.Printf("Membagi %d dengan %d\n", a, b)
	hasil := a / b //ini bisa sebabkan panic kalau b = 0
	fmt.Println("Hasil:", hasil)
}
