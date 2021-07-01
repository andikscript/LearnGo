package main

import "fmt"

// deklarasi variable di luar fungsi tidak akan terjadi eror
// jika variable tersebut tidak digunakan
var a = 12
// deklarasi inference diluar fungsi tidak dapat di lakukan atau akan terjadi eror
// b := 11

func main() {
	// semua deklarasi variable di dalam fungsi semua harus dipakai, jika tidak maka akan terjadi eror

	// deklarasi variable beserta tipe data (manifest typing)
	var firstName string = "andik"
	var lastName string
	lastName = "script"

	fmt.Println("first name :", firstName, "last name :", lastName)

	// deklarasi variable tanpa tipe data (type inference)
	alamat := "ponorogo"
	var status = "mahasiswa"

	fmt.Println(alamat)
	fmt.Println(status)

	// deklarasi multi variable
	var a, b, c string
	a, b, c = "a", "b", "c"
	fmt.Println(a, b, c)

	var d, e, ee = 12, 13, "ee"
	fmt.Println(d, e, ee)

	f, g, h := "f", "g", 12.4
	fmt.Println(f, g, h)

	// deklarasi variable _ (underscore) atau reserved variable
	// variable yang digunakan untuk menampung nilai yang tidak digunakan (keranjang sampah)
	// nilai yang ditampung di dalam variable _ tidak bisa ditampilkan
	var _ = 12
	r, _ := "as", 12

	fmt.Println(r)

	// deklarasi variable pointer
	nilai := new(string)
	nilai = &r
	fmt.Println(nilai)
	fmt.Println(*nilai)
}
