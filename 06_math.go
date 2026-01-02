package main

import "fmt"

func math() {
	fmt.Println("==============ARITHMETIC OPERATORS============")
	// operasi dasar utk lakukan perhitungan pda angka
	// bisa di tipe data numerik-> int, float32, float64
	// +, -, *, /, %

	a := 10
	b := 3
	fmt.Println("a + b:", a+b)
	fmt.Println("a - b:", a-b)
	fmt.Println("a * b:", a*b)
	fmt.Println("a / b:", a/b)
	fmt.Println("a % b:", a%b)

	c := 7.8
	d := 2.3
	fmt.Println("c / d =", c/d)

	fmt.Println("==============ASSIGNMENT OPERATORS============")
	x := 10
	x = 5

	x += 2
	fmt.Println("x +=2 :", x)
	x -= 2
	fmt.Println("x -=2 :", x)
	x *= 2
	fmt.Println("x *=2 :", x)
	x /= 2
	fmt.Println("x /=2 :", x)
	x %= 2
	fmt.Println("x %=2 :", x)

	fmt.Println("==============INCREMENT & DECREMENT============")
	y := 20
	fmt.Println(y)
	y++
	fmt.Println(y)
	y--
	fmt.Println(y)

	fmt.Println("==============COMPARISON OPERATORS============")
	// utk bandingkan dua nilai dan hasilnya selalu bool (true/false)
	// bisa utk tipe data numerik, string, dan bool
	// tdk bisa utk slice, map, function
	// sering dipake di kondisional cth if-else, switch
	// ==, !=, >, <, >=, <=

	var angka1, angka2 int
	// input dari user
	fmt.Println("Masukkan angka pertama:")
	fmt.Scanln(&angka1)
	fmt.Println("Masukkan angka kedua:")
	fmt.Scanln(&angka2)

	fmt.Println("\n<==== Hasil Perbandingan ====>")
	// %d-> placeholder/string format utk int (decimal)
	// %v-> hasil dari nilai apapun (generic)
	fmt.Printf("%d == %d ? %v\n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d != %d ? %v\n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d > %d ? %v\n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d < %d ? %v\n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d >= %d ? %v\n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d <= %d ? %v\n", angka1, angka2, angka1 < angka2)

	fmt.Println("==============LOGIC OPERATORS============")
	// utk gabungkan 2/lebih ekspresi boolean (benar/salah)
	// biasanya pake if, for dll
	// &&-> AND, ||-> OR, !-> NOT
	fmt.Println("AND: ", (5 > 3) && (8 > 6)) //true
	fmt.Println("AND: ", (5 > 3) && (8 < 6)) //false

	fmt.Println("OR: ", (5 > 3) || (8 < 6)) //true
	fmt.Println("OR: ", (5 < 3) || (8 < 6)) //false

	fmt.Println("NOT: ", !(5 > 3)) //false
	fmt.Println("NOT: ", !(5 < 3)) //true
}
