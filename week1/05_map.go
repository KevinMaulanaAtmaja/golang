package week1

import "fmt"

func Maps() {
	fmt.Println("=============MAP=============")
	// map=> stuktur data utk simpan psangan key-value
	// ibarat-> kamus, cari kata(key), utk dpt artinya(value)

	// key harus tipe data yg bisa dibandingkan, cth string, int, bool, float dll
	// unique key, value bisa tipe data apapun cth string, int, bool, slice
	// sifat unordered, tdk ada urutan ttap
	// akses cepat/O(1), pake hash table jadinya cepat
	// analogi: mirip loker pribadi 1 kunci = 1 isi

	// pake literal
	m := map[string]int{
		"apel":   10,
		"jeruk":  20,
		"pisang": 30,
	}
	fmt.Println(m) //map[apel:10 jeruk:20 pisang:30]

	// pake make
	m2 := make(map[string]int)
	m2["anggur"] = 40
	m2["pir"] = 50
	fmt.Println(m2) //map[anggur:40 pir:50]

	// map kosong
	m3 := make(map[string]int, 0)
	fmt.Println(m3)

	// map dg data
	m4 := map[string]int{"a": 1}
	fmt.Println(m4)

	fmt.Println("=============MAP FUNCTION=============")
	mhs := map[string]string{
		"nama":  "Kevin",
		"prodi": "A-TRPL",
	}

	// tmbh/update key
	mhs["nama"] = "Kevin Maulana"
	mhs["jurusan"] = "JBI"

	// ambil value
	val := mhs["prodi"]
	fmt.Println(val)

	fmt.Println(mhs)
	fmt.Println("Nama:", mhs["nama"], "Prodi:", mhs["prodi"], "Jurusan:", mhs["jurusan"])

	// cek apkh key ada
	val, ok := mhs["prodi"]
	fmt.Println(val, ok)

	// panjang map
	fmt.Println("Panjang Map:", len(mhs))

	// hapus key
	delete(mhs, "prodi")
	fmt.Println(mhs)
}
