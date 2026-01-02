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
}
