package defpanrec

import "fmt"

func Defpanrec() {
	fmt.Println("=============DEFPANREC EXAMPLE=============")
	// pake defer utk cleanup-> memastikan resource slalu dirilis
	// jngn pake defer di loop besar, boros. taruh defer di func terpisah
	// panic hanya utk kondisi datal, biasakan return error utk kasus yg dpt diprediksi
	// recover utk buat boundary, agar app tdk langsung crash

	// cth fungsi cleanup sederhana
	// bacaConfig("ada_nama_filenya")
	bacaConfig("")
	fmt.Print("Program tetap berjalan setelah panic")
}

func cleanup() {
	fmt.Println("Cleanup: menutup resource...")
}

func bacaConfig(namaFile string) {
	// defer slalu dipanggil meskipun ada panic
	defer cleanup()

	// recover diletakkan dlm defer utk tangkap panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error:", r)
		}
	}()

	if namaFile == "" {
		panic("Nama file konfigurasi tidak boleh kosong")
	}

	fmt.Println("Membaca konfigurasi dari", namaFile)
	// anggap berhasil membaca
}
