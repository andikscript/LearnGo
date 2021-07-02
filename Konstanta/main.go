package main

import "fmt"

func main() {
	// deklarasi konstanta harus sekaligus diisi nilainya
	// jika tidak akan mengalami eror dan konstanta
	// tidak akan mengalami eror jika suatu konstanta
	// tidak digunakan atau dipakai
	// dan deklarasi konstanta tidak perlu di deklarasikan tipe konstantanya
	// karena konstanta menggunakan model deklarasi type inference meskipun sebenarnya bisa
	// menggunakan tipe data
	const nilai = 19

	// konstanta tidak bisa diubah nilainya
	fmt.Println(nilai)

	const key string = "asdasd123xasdasasdf'l'asda"
}
