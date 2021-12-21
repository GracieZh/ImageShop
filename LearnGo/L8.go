package main
import (
	"fmt"
	"time"
	"sync"
)

func try1() {
	defer fmt.Println("Defer print")
	go sayHello()
	time.Sleep(0*time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func try2() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg *string) {
		fmt.Println(*msg)
		wg.Done()
	}(&msg)
	msg = "Goodbye"
	wg.Wait()
}

//////////////////////////////////////////////

var counter = 0

func SayHello() {
	fmt.Println("Hello # ", counter)
	m.RUnlock()
	wg.Done()
}

func Increment() {
	counter ++
	m.Unlock()
	wg.Done()
}

func try3() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go SayHello()
		m.Lock()
		go Increment()
	}
	//fmt.Println("1 counter = ", counter)	
	wg.Wait()
	fmt.Println("2 counter = ", counter)	
}

//////////////////////////////////////////////

func main() {
	i := 3
	switch i {
	case 1: try1()
	case 2: try2()
	case 3: try3()
	default: return
	}
		
}

