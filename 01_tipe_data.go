package main

import (
	"fmt"
	"strings"
)

func dataType() {
	fmt.Println("=============VARIABLE=============")
	// variable => wadah/kotak utk simpan data
	// cara klasik: sebutkan tipe datanya
	// titik koma optional
	var name string = "Kevin"
	var age int = 19

	// golang juga bisa otomatis menebak tipe datanya (type inference)
	city := "Banyuwangi"
	year := 2026

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Year:", year)

	fmt.Println("=============KONSTANTA=============")
	// konstanta => nilai tak bisa diubah
	// utk data yg sifatnya tetap, cth nama app, rumus, dll
	const pi = 3.14
	const appName = "Belajar Golang"

	fmt.Println("Nilai Pi:", pi)
	fmt.Println("App Name:", appName)

	// jika coba diubah akan error
	// pi = 3.1415

	fmt.Println("=============NUMBER=============")
	// tipe data number
	// integer-> angka bulat
	// floating point-> angka desimal

	// signed integers
	var i8 int8 = 127
	var i16 int16 = -32767
	var i32 int32 = 2147483647
	var i64 int64 = -9223372036854775807
	var i int = -100

	// unsigned integers
	var u8 uint8 = 127
	var u16 uint16 = 32767
	var u32 uint32 = 2147483647
	var u64 uint64 = 9223372036854775807
	var u uint = 100

	// printf-> utk print dg format
	fmt.Println("Signed integers: ")
	fmt.Printf("int8   : %v\n", i8)
	fmt.Printf("int16  : %v\n", i16)
	fmt.Printf("int32  : %v\n", i32)
	fmt.Printf("int64  : %v\n", i64)
	fmt.Printf("int    : %v\n", i)

	fmt.Println("\nUnsigned integers: ")
	fmt.Printf("uint8   : %v\n", u8)
	fmt.Printf("uint16  : %v\n", u16)
	fmt.Printf("uint32  : %v\n", u32)
	fmt.Printf("uint64  : %v\n", u64)
	fmt.Printf("uint    : %v\n", u)

	fmt.Println("=============FLOAT=============")
	// floating point
	var f32 float32 = 3.1415926535
	var f64 float64 = 3.14159265358979323846

	fmt.Println("Nilai float32:", f32)
	fmt.Println("Nilai float64:", f64)

	fmt.Println("=============BOOL=============")
	var isRaining bool = true
	var isSunny bool = false
	isLogedIn := true
	hasPermission := false

	fmt.Println("Apakah hujan?", isRaining)
	fmt.Println("Apakah cerah?", isSunny)
	fmt.Println("Apakah login?", isLogedIn)
	fmt.Println("Apakah punya akses?", hasPermission)

	fmt.Println("=============STRING=============")
	// string => kumpulan karakter yg ditulis diantara tanda kutip ganda("")/backtick(`)
	// utk simpan dan manipulasi teks, imutable, slice of bytes
	var nama string = "Amel"
	pesan := "Amel Cantik!"

	paragraf := `Haloo ini adalah contoh 
	multi-line string di Go.`

	fmt.Println("Nama:", nama)
	fmt.Println("Pesan:", pesan)
	fmt.Println("Paragraf:\n", paragraf)

	fmt.Println("=============MANIPULASI STRING=============")
	// manipulasi string
	text := "Belajar Golang"

	text2 := "Halo" + ", " + "Golang" + "!"
	fmt.Println(text2)
	fmt.Println("Length:", len(text))
	fmt.Println("Go"[0])
	fmt.Println(string("Go"[0]))
	// text3 := "Go"[0] = 'g' -> error imutable

	fmt.Println("Lowercase:", strings.ToLower(text))                       // belajar golang
	fmt.Println("Uppercase:", strings.ToUpper(text))                       // BELAJAR GOLANG
	fmt.Println("StartsWith Belajar?", strings.HasPrefix(text, "Belajar")) //true
	fmt.Println("Contains Golang?", strings.Contains(text, "Golang"))      //true

	parts := strings.Split("apel,banana,cerry", ",") //delimiter
	fmt.Println("Split:", parts)                     //[apel banana cerry]

	newText := strings.ReplaceAll(text, "Belajar", "Halo")
	fmt.Println("Replace:", newText) // Halo Golang!
	fmt.Println("Panjang teks:", len(text))
}
