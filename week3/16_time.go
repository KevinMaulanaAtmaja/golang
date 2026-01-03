package week3

import (
	"fmt"
	"time"
)

func Times() {
	fmt.Println("=============TIMEZONE=============")
	// tipe data waktu dan tgl disediakan oleh package time
	// time-> simpan zona waktu (default ikuti sistem/os)
	// zona waktu mempengaruhi cara waktu ditmpilkan, tpi nilai UTC-nya ttap sama,
	// UTC-> waktu universal/standar waktu acuan global, local-> zona waktu lokal(ikuti os/server), custom location-> zona waktu ttntu sperti "Asia/Jakarta"
	// UTC tdk brubah karna DST(Daylight Saving Time) dipake utk mempercepat waktu di negara yg memiliki DST
	// UTC utk leap seconds sskali agar sinkron dg rotasi bumi, jarang relevan utk aplikasi biasa, cocok utk komunikasi antar sistem & penyimpanan

	// pake zone waktu trtntu
	// time.LoadLocation("Asia/Jakarta")-> memuat lokasi dari db IANA Time Zone
	// Asia/Jakarta, Asia/Makassar, Asia/Jayapura, UTC
	// loc, _ := time.LoadLocation("Asia/Jakarta")
	// fmt.Println("Waktu Jakarta: ", now.In(loc))
	// server vs local-> jika server luar negeri, time.Now() akan ikut zona waktu server
	// simpan waktu di db dlm format UTC, lalu konversi ke zona waktu user saat ditampilkan

	now := time.Now() //ambil waktu saat ini
	fmt.Println("Sekarang: ", now)

	//1. ambil waktu saat ini dlm utc
	nowUTC := time.Now().UTC()
	fmt.Println("Waktu UTC:", nowUTC)

	// 2. konversi ke zona waktu lokal (Asia/Jakarta)
	jakartaTime := convertToZone(nowUTC, "Asia/Jakarta")
	fmt.Println("Waktu Jakarta:", jakartaTime)

	// 3. konversi ke zona waktu lain yg pake DST (misal new york)
	newYorkTime := convertToZone(nowUTC, "America/New_York")
	fmt.Println("Waktu New York:", newYorkTime)

	// 4. simpan waktu (simulasi penyimpanan ke db)
	// disarankan simpan waktu dlm utc
	fmt.Println("Disimpan ke database (UTC):", nowUTC)
}

func convertToZone(t time.Time, zone string) time.Time {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		fmt.Println("Error memuat lokasi:", err)
		return t
	}
	return t.In(loc)
}
