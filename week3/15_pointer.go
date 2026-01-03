package week3

import "fmt"

func Pointers() {
	fmt.Println("=============POINTER=============")
	// sbuah variabel yg simpan alamat memo dari nilai variable lain
	// bukan hanya nilai lngsung yg disimpan, tpi lokasi tmpt nilai itu brda di memo
	// pointer = penunjuk alamat memo, var biasa == simpan nilai langsung
	// pointer mungkinkan kita modifikasi variable asli tanpa harus kembalikan dari fungsi
	// simbol & utk ambil alamat memo, simbol * utk akses isi dari alamat itu

	// cth
	x := 10
	p := &x
	// 	&x → ambil alamat x
	// p simpan alamat itu
	// *p = 20 → ubah isi alamat
	// x ikut berubah
	*p = 20
	fmt.Println(x) // 20

	// ubah nama
	nama := "Amel"
	ubahNama(&nama)   //kirim alamat var nama
	fmt.Println(nama) //output: Kevin(berubah!)

	// game
	hp := 50 //hp awal
	fmt.Println("HP Awal:", hp)

	heal(&hp) //kirim alamat hp
	fmt.Println("HP Sekarang:", hp)

	attack(&hp, 30) //kirim alamat hp
	fmt.Println("HP Sekarang:", hp)

	attack(&hp, 40) //ini bikin hp <= 0
	fmt.Println("HP Sekarang:", hp)

}
func ubahNama(nama *string) {
	*nama = "Kevin" //ubah lngsung ke alamat memo aslinya
}

func heal(hp *int) {
	*hp += 20 // tmbh hp lngsung di memo
	fmt.Println("Permain disembuhkan +20!")
}

func attack(hp *int, damage int) {
	*hp -= damage
	fmt.Printf("Permain diserang! HP berkurang %d!\n", damage)
	if *hp <= 0 {
		fmt.Println("Game Over! Pemain kalah 😵")
	}
}
