package main

import "fmt"

func slice() {
	fmt.Println("============SLICE==============")
	// slice=> array yg bisa ditmbh/dikurangi, ukuran bisa berubah
	// komponen => pointer, length, dan capacity
	// pointer-> menunjuk ke elemen awal di array (bisa ditngah/diakhir, fleksibel)
	// length-> brp bnyk elemen yg terlihat di slice
	// capacity-> brp bnyk elemen dari pointer smpai akhir array
	// array(panjangt ttap, elemen tersimpan lngsung di memo)
	//  [10, 20, 30, 40, 50]
	//   0,  1,  2,  3,  4 -> indeks array

	// slice(punya pointer, length, capacity)
	// s := arr[start:end]

	// slice: [20, 30, 40]
	// pointer -> arr[1] nilai 20
	// length -> 3 (20, 30, 40)
	// capacity -> 4 (20, 30, 40, 50)

	arr := [5]int{1, 2, 3, 4, 5}
	// slice sluruh array
	s1 := arr[:]
	fmt.Println("Slice s1:", s1)

	// slice dari index awal ke trtntu
	s2 := arr[:3]
	fmt.Println("Slice s2:", s2)

	// slice dari index trtntu ke akhir
	s3 := arr[2:]
	fmt.Println("Slice s3:", s3)

	// slice dari index i ke j
	s4 := arr[1:4]
	fmt.Println("Slice s4:", s4)

	fmt.Println("============SLICE FUNCTION==============")
	s5 := make([]int, 3, 5) //slice tipe int, length 3, capacity 5
	fmt.Println(s5)
	fmt.Println("Len:", len(s5), "cap:", cap(s5)) // len 3, cap 5

	s6 := append(s5, 10, 80)     // [0, 0, 0, 10, 80] -> len(5), cap(5)
	s7 := append(s5, 10, 80, 90) // [0, 0, 0, 10, 90]
	fmt.Println(s6)
	fmt.Println("Len Setelah di append:", len(s6))
	fmt.Println("Len:", len(s7), "cap:", cap(s7)) // nmbh otomatis cap klo lebih dari yg ditntukan-> len(6), cap(10)

	slice2 := make([]int, 3)
	n := copy(slice2, s6) //copy(dst, src)
	fmt.Println("Copied:", slice2)
	fmt.Println("Jumlah elemen yg tersalin:", n)

	angka := []int{1, 2, 3, 4, 5}
	slice4 := angka[1:4]
	fmt.Println("Slice 4:", slice4)
}
