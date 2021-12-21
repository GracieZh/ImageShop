package main

import ("fmt")

func tryShift() {
	var v1 int64  = 1
	v2 := v1 << 3
	v3 := 0xc0
	v4 := v3 >> 5
	fmt.Printf("%v %b %X \n", v1, v1, v1)	
	fmt.Printf("%v %b %X \n", v2, v2, v2)	
	fmt.Printf("%v %b %X \n", v3, v3, v3)	
	fmt.Printf("%v %b %X \n", v4, v4, v4)	
}

func tryArray() {
	v1 := [3]int{1,2}
	v2 := [...]int{1,2,3,4,5}
	var v3 [3]int
	v4 := v2
	fmt.Printf("%v, %T \n", v1, v1)
	fmt.Printf("%v, %T \n", v2, v2)
	fmt.Printf("%v, %T \n", v3, v3)
	fmt.Printf("%v, %T, %v \n", v4, v4, len(v4))
}

func tryArray2() {

	var m [3][3]int = [3][3]int { [3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1} }
	fmt.Println(m)

	a := []int{1,2,3,4,5,6,7,8,9,10}
	b := a[:] // slice of all elements
	c := a[3:] // slice from 4th element to end
	d := a[:6] // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th, and 6th elements

	f := &a
	(*f)[1] = 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%v %T \n", (*f)[1], (*f)[1])
}

func tryArray3() {

	v1:=make([]int,2) // create slice with capacity and lenght == ?
	v2:=make([]int,2,10) // slice with length == ? and capacity == ??
	v3 := 1
	v4 := 1
	v2 = []int{1,2,3}
	//v1 = []int{1,2,3}


	fmt.Printf("%v, %T \n", v1, v1)
	fmt.Printf("%v, %T \n", v2, v2)
	fmt.Printf("%v, %T \n", v3, v3)
	fmt.Printf("%v, %T \n", v4, v4)
}


func main() {
	i:=3
	if i==1 { tryShift() }
	if (i==2) { tryArray() }
	if (i==3) { tryArray2() }
}
