package week3

import "fmt"

func Structs() {
	fmt.Println("=============STRUCT=============")
	// struct(structure)=> utk kelompokkan bbrp data dg tipe yg beda dlm satu kesatuan
	// kerangka/blueprint utk buat abj dg properti trtntu
	// Go tdk punya class & inheritance, tpi punya struct + interface
	// struct tdk punya inheritance, tpi bisa pake composition(sisipkan struct lain di dlm struct)
	// struct sering dipake ama method supaya bisa punya aksi (mirip OOP)
	// method-> fungsi yg nempel di struct, utk beri perilaku(behavior) ke sbuah strict, mirip OOP
	// fungsi struct:
	// 1.klompokkan data yg berkaitan, cth User: name, email, age
	// 2.membuat representasi obj nyata, cth: Mobil, Mhs, Buku
	// 3.permudah modularisasi kode agar lbih rapi, terstruktur, dan mudah dibaca

	// panggil obj
	user1 := User{
		Name:  "Kevin",
		Email: "kevin@gmail.com",
		Age:   19,
	}
	user2 := User2{
		Name:  "Amel",
		Email: "amel@gmail.com",
		Address: Address{
			City:     "Banyuwangi",
			Province: "East Java",
		},
	}
	fmt.Println("Name:", user1.Name, "Email:", user1.Email, "Age:", user1.Age)
	fmt.Println("Name:", user2.Name, "Email:", user2.Email, "Province:", user2.Address.Province, "City:", user2.Address.City)

	pp := PersegiPanjang{Panjang: 10, Lebar: 5}
	pp2 := PersegiPanjang{Panjang: 4, Lebar: 3}
	fmt.Println("Panjang:", pp.Panjang, "Lebar:", pp.Lebar)
	fmt.Println("Panjang:", pp2.Panjang, "Lebar:", pp2.Lebar)
	fmt.Println("Luas pp:", pp.luas())
	fmt.Println("Luas pp2:", pp2.luas())
	fmt.Println("Luas Biasa:", LuasBiasa(2, 3))
}

// cara bikin struct
type User struct {
	Name  string
	Email string
	Age   int
}
type User2 struct {
	Name    string
	Email   string
	Address //embedded struct (composition)
}

// struct lain
type Address struct {
	City, Province string
}

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

func (p PersegiPanjang) luas() float64 {
	return p.Panjang * p.Lebar
}

func LuasBiasa(panjang, lebar int) int {
	return panjang * lebar
}
