// Variable declaration
package main

import "fmt"

var i float32 = 42

func tryInt() {
	var i1 int8
	var i2 uint32 = 42
	i3 := -100
	b1 := true
	c1 := "Hello bear!"
	c2 := 'a'
	var c4 uint8 = 'b'
	c3 := string(c4)
	fmt.Printf("%v, %T \n", i1, i1)
	fmt.Printf("%v, %T \n", i2, i2)
	fmt.Printf("%v, %T \n", i3, i3)
	fmt.Printf("%v, %T \n", b1, b1)
	fmt.Printf("%v, %T \n", c1, c1)
	fmt.Printf("%v, %T \n", c2, c2)
	fmt.Printf("%v, %T \n", c3, c3)
	fmt.Printf("%v, %T \n", c4, c4)
}

func tryFloat() {
	var v1 float64
	var v2 float32 = 3.14
	v3 := -15E18
	v4 := 16.7e12
	fmt.Printf("%v, %T \n", v1, v1)
	fmt.Printf("%v, %T \n", v2, v2)
	fmt.Printf("%v, %T \n", v3, v3)
	fmt.Printf("%v, %T \n", v4, v4)
}

func tryComplex() {
	var v1 complex64
	var v2 complex64 = 3.14 + 2i
	v3 := 3 - 4i
	v4 := v2 + complex64(v3)
	fmt.Printf("%v, %T \n", v1, v1)
	fmt.Printf("%v, %T \n", v2, v2)
	fmt.Printf("%v, %T \n", v3, v3)
	fmt.Printf("%v, %T \n", v4, v4)
	fmt.Printf("%v + %v i \n", real(v4), imag(v4))
}

func tryText() {
	var v1 string
	var v2 string = "abcdefg"
	var v uint8 = 'u'
	v3 := v2 + " " + string(v)
	v4 := []byte(v2)
	var c1 rune ='a'
	c2 := 'П'
	c3 := "Привіт Світ "
	c4 := c3 + string(c2) + " " + v3
	c5 := []byte(c4)
	c6 := string(c5)
	c7 := []rune(c6)
	c8 := string(c7)

	fmt.Printf("%v, %T \n", v1, v1)
	fmt.Printf("%v, %T \n", v2, v2)
	fmt.Printf("%v, %T \n", v, v)
	fmt.Printf("%v, %T \n", v3, v3)
	fmt.Printf("%v, %T \n", v4, v4)
	fmt.Printf("%v, %T \n", c1, c1)
	fmt.Printf("%v, %T \n", c2, c2)
	fmt.Printf("%v, %T \n", c3, c3)
	fmt.Printf("%v, %T \n", c4, c4)
	fmt.Printf("%v, %T \n", c5, c5)
	fmt.Printf("%v, %T \n", c6, c6)
	fmt.Printf("%v, %T \n", c7, c7)
	fmt.Printf("%v, %T \n", c8, c8)
}

func tryConst() {

	const (
		a = iota
		b
		c
	)
	fmt.Printf("%v %v %v\n", a, b, c)

	const v1 float64 = 3.14
	const v2 int = 14
	const v3 string = "foo"
	const v4 bool = true
	fmt.Printf("%v, %T \n", v1, v1)
	fmt.Printf("%v, %T \n", v2, v2)
	fmt.Printf("%v, %T \n", v3, v3)
	fmt.Printf("%v, %T \n", v4, v4)
	
	const (
		_ = iota // ignore first value by assigning to blank identifier
		KB = 1 << ( 10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fileSize := 4000000000.

	fmt.Printf("%.3f GB\n", fileSize/GB)



}

func main() {
	i:=4
	if i==1 { tryInt() }
	if (i==2) { tryFloat() }
	if (i==3) { tryComplex() }
	if (i==4) { tryText() }
	if (i==5) { tryConst() }
}
