package main

import "fmt"

func main() {
	// tipe data numerik non desimal
	var a uint8
	a = 255
	var b byte = 255
	var bb uint = 213131231

	fmt.Printf("%d \n", a)
	fmt.Printf("%d \n", b)
	fmt.Printf("%d \n", bb)

	var c int8
	c = -127
	var d int16 = -32768
	var f rune
	f = 1212121
	var ff = -323232 // defaultnya int32 / int

	fmt.Printf("%d \n", c)
	fmt.Printf("%d \n", d)
	fmt.Printf("%d \n", f)
	fmt.Printf("%d \n", ff)

	// tipe data numerik desimal
	var g float32
	g = 1212.34
	var h float64 = 453452.23423

	// default nol dibelakang titik ada 6
	fmt.Printf("%f \n", g)
	fmt.Printf("%.4f \n", h)
	// perbedaan float32 dengan float64 adalah cakupannya

	// tipe data boolean
	var i bool
	i = true
	var j bool = false

	fmt.Printf("%t \n", i)
	fmt.Printf("%t \n", j)

	// tipe data string
	var k string
	k = "ini string"
	var l string = "ini juga string"

	fmt.Printf("%s \n", k)
	fmt.Printf("%s \n", l)

	var m = `ini adalah penggunaan tanda accent
semua tanda %s \n "" dan simbol yang lain tidak akan dianggap istimewa`

	fmt.Printf("%s \n", m)

	var n = map[int]interface{}{1:12,2:23,3:34}
	fmt.Println(n[1].(int) * 12)



	var sl = []int{10,20,30}
	var ma = map[*int]*int{}

	for key, val := range sl {
		ma[&key] = &val
	}

	for i2, i3 := range ma {
		fmt.Println(*i2, ":",*i3)
	}

	var t = [2]string{}
	t[0] = "as"
	t[1] = "as"

	fmt.Println(t)
}
