package main

import "fmt"

func main() {
	// operator aritmatika
	// untuk operasi angka
	var a, b = 12, 15

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// operator pembanding
	// menghasilkan nilai true atau false

	fmt.Println(a == b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)

	// operator logika
	// untuk operasi variable yang bernilai true atau false
	var c, d = true, true

	fmt.Println(c && d)
	fmt.Println(c || d)
	fmt.Println(!c)
}
