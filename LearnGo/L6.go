package main

import "fmt"

func tryLoop1() {
	for i,j:=1,1; i<5; i,j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println("")

	for i:=1; i<5; i++ {
		fmt.Println(i)
	}
}

func tryLoop2() {
	m1 := []int{1,2,3}
	for index, value := range m1 {
		fmt.Println(index, value)
	}
	for _, value := range m1 {
		fmt.Println(value)
	}
	s := "Привіт Світ"
	for index, value := range s {
		fmt.Println(index, string(value))
	}

}

func tryLoop3() {
	m1 := map[string]int{
		"fox": 1,
		"frog": 2,
		"ant": 3,
		"lion": 4,	
		"fish": 5,
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

func tryFunc4() {
	f1 := func (a int, b int) { 
		a = 11
		b = 21
	}
	f2 := func (a *int, b *int) { 
		*a = 11
		*b = 21
	}
	a := 1
	b := 2

	fmt.Println(a, b);
	f1(a,b)
	fmt.Println(a, b);
	f2(&a, &b)
	fmt.Println(a, b);
}

func tryFunc5() {
	sum := func (msg string, values ...int) (int, error) { 
		
		if len(values) > 5 { 
			return 0, fmt.Errorf("Too many values") 
		}

		fmt.Println(values)
		result := 0
		for _,v := range values {
			result += v
		}
		fmt.Println(msg, result)
		return result, nil
	}

	res, err := sum("The sum is", 1,2,3,4,5)
	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Println("result is:", res)
	}

	res, err = sum("The sum is", 1,2,3,4,5,6,7)
	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Println("result is:", res)
	}

}

func tryFunc6() {
	sum := func (values ...int) (result int) { 
		result = 0
		for _,v := range values {
			result += v
		}
		return
	}

	res := sum(1,2,3,4,5)
	fmt.Println("result is:", res)

	product := func (values ...int) *int { 
		result := 1
		for _,v := range values {
			result *= v
		}
		return &result
	}

	prod := product(1,2,3,4)
	fmt.Println("result is:", *prod)

}

func main () {
	i := 5
	switch i {
	case 1: tryLoop1()
	case 2: tryLoop2()
	case 3: tryLoop3()
	case 4: tryFunc4()
	case 5: tryFunc5()
	case 6: tryFunc6()
	default: return
	}
}