package week2

import "fmt"

func Loop() {
	fmt.Println("=============LOOP=============")
	// for satu-satunya loop yg tersedia
	// penggunaannya fleksibel bisa gantikan while/do while

	// stuctur
	// for inisialisasi; kondisi; post {
	// blok kode yg akan diulang
	// }
	// inisialisasi-> utk deklarasi variable counter
	// kondisi-> ekspresi boolean, jika true loop jalan
	// post-> ekspresi yg dieksekusi setiap iterasi(i++/i--)

	// for standar
	fmt.Println("\n1. for standar (i = 1; i <= 5; i++):")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// while-style
	fmt.Println("\n2. while-style (hanya kondisi):")
	j := 2
	for j <= 10 {
		fmt.Printf("%d  ", j)
		j += 2 // j = j+2
	}
	fmt.Println()

	// infinite loop
	fmt.Println("\n3. infinite loop dg break:")
	k := 5
	for {
		fmt.Printf("%d ", k)
		k-- // k = k-1
		if k == 0 {
			break //keluar dari loop
		}
	}
	fmt.Println()

	// range loop
	fmt.Println("\n4. range loop:")
	buah := []string{"apel", "jeruk", "pisang"}
	for index, value := range buah {
		fmt.Printf("Index %d: %s\n", index, value)
	}
	fmt.Println()

	// for dg break dan continue
	fmt.Println("\n5. for dg break dan continue:")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue //skip angka 3
		}
		if i == 5 {
			break //keluar di angka 5
		}
		fmt.Println(i)
	}
}
