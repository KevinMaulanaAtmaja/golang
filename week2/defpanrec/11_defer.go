package defpanrec

import "fmt"

func Defers() {
	fmt.Println("=============DEFER=============")
	// defer=> fitur utk jadwalkan sbuah fungsi  utk dipanggil trkhir kali sblum fungsi utama selesai
	// akan jalan dulu meski terjdi "panic",
	// jika ada bbrp defer maka LIFO(last in first out)
	// argumen dari fungsi yg didefer dieval saat baris defer ditulis, bukan saat dieksekusi, artinya...
	// jika x = 1, lalu diubh jdi 2 sblum defer maka yg keluar ttap 1
	// x := 10
	// defer fmt.Println(x) //10 bukan 20
	// x = 20

	// defer cth
	defer fmt.Println("A: dieksekusi paling akhir (LIFO)")
	defer fmt.Println("B: dieksekusi agak akhir")
	fmt.Println("C: dieksekusi duluan")

	x := 1
	defer fmt.Println("defer-1, x =", x) // argumen x dievaluasi skrng(1)
	x = 2
	defer fmt.Println("defer-2, x =", x) // argumen x dievaluasi skrng(2)
	fmt.Println("body, x =", x)
}
