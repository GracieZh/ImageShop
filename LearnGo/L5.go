package main
import (
	"fmt"
	"reflect"
)

func tryMap() {
	m1 := map[string]int{
		"fox": 1,
		"flog": 2,
		"ant": 3,
		"lion": 4,
		"fish": 5,
	}
	m2 := map[int]string{}

	m3 := make(map[string]int)
	m3 = m1

	m3["toad"]=1


	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
}

type Animal struct {
	Name string `required max:"100"`
	Origin string
}

func tryStruct() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

func tryStruct2() {
	type Doctor struct {
		number int
		actorName string
		companions []string
	}
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
}

func tryStruct3() {
	aDoctor := struct {
		number int
		name string
		companions []string
	} {
		number: 3,
		name: "Jon Snow",
		companions: []string{
			"Liz Shaw",
			"Bob Rainbow",
			"Sarah Jane Smith",
		},
	}
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"
	anotherDoctor.companions = []string{"Bread","Pancake","Cookie","Icecream Sandwich"}
	fmt.Println(aDoctor)
	fmt.Println(anotherDoctor)
}

func tryStruct5() {
	type Animal struct {
		name string
		speed float32
	}

	type Bird struct {
		Animal
		canFly bool
	}
	
	sparrow := Bird {
		Animal: Animal { name: "Amy", speed: 50, },
		canFly: true,
	}

	/*
	sparrow := Bird{}
	sparrow.name = "Amy"
	sparrow.speed = 50
	sparrow.canFly = true
	*/
	
	fmt.Println(sparrow)
}

func trySwitch6() {
	
	switch i:=2+3; i {
	case 1,5,10:
		fmt.Println("one five ten")
	case 2,4,6:
		fmt.Println("two four six")
	default:
		fmt.Println("another number")
	}
}

func trySwitch7() {
	i:=11
	switch {
	case i<=10:
		fmt.Println("less or equals to 10")
	case i<=20:
		fmt.Println("less or equals to 20")
	default:
		fmt.Println("else")
	}
}

func trySwitch8() {
	//var i interface{} = 3.14 // [3]int{}
	//switch i.(type) {
	switch interface{}(3.14).(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is a [3]int")
	default:
		fmt.Println("i is another type")
	}
}

func main() {
	i := 8
	if i==1 { tryMap() 
	} else if i==2 { tryStruct() 
	} else if i==3 { tryStruct2() 
	} else if i==4 { tryStruct3() 
	} else if i==5 { tryStruct5() 
	} else if i==6 { trySwitch6() 
	} else if i==7 { trySwitch7() 
	} else if i==8 { trySwitch8() 
	}
}