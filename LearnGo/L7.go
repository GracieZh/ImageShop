package main
import ("fmt")

type Animal struct {
	name string
	speed int
}

func (a Animal) greet() {
	fmt.Println("Hello from", a.name)
}

func (a *Animal) goodbye(name string) {
	fmt.Println(a.name, "says: Bye", name)
}

func tryFunc1() {
	fish := Animal { name: "Amy"}
	bird := Animal { name: "Bob"}
	fish.greet()
	bird.greet()
	(&fish).goodbye(bird.name)
	(&bird).goodbye(fish.name)
}

//////////////////////////////////////

type Speaker interface {
	Speak([]byte) (int, error)
}
type Mover interface {
	Move(int) int
}

type Pet interface {
	Speaker
	Mover
}

type Cat struct {}

func (cw Cat) Speak(data []byte) (int, error) {
	n, err := fmt.Println("Meow, ",string(data))
	return n, err
}

func (cw Cat) Move(distance int) int {
	n, _ := fmt.Println("Meow, I walks ",distance, " cm ")
	return n
}

type Fish struct {}

func (cw Fish) Speak(data []byte) (int, error) {
	n, err := fmt.Println("Fish says", string(data))
	return n, err
}

func (cw Fish) Move(distance int) int {
	n, _ := fmt.Println("Fish swims ",distance, " cm ")
	return n
}

func tryFunc2() {
	var w interface{} = Cat{}
	if pet, ok := w.(Pet); ok {
		pet.Speak([]byte("Hello Go!"))
	}

	var cat Pet = Cat{}
	cat.Speak([]byte("Hello W1!"))

	var fish Pet = Fish{}
	fish.Speak([]byte("Hello W2!"))
	cat.Move(100)
	fish.Move(50)

}

//////////////////////////////////////

type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic ++
	return int(*ic)
}

func tryFunc3() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i<10; i++ {
		fmt.Println(inc.Increment())
	}
}

//////////////////////////////////////

func main() {
	i:=2
	switch i {
	case 1: tryFunc1()
	case 2: tryFunc2()
	case 3: tryFunc3()
	default: return
	}
}