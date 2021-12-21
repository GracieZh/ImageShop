package main
import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

var x = 111

func tryChannel1() {
	ch := make(chan int)
	for j := 0; j<5; j++ {
		wg.Add(2)
		go func() {
			i := <- ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			ch <- x
			x += 100
			wg.Done()
		}()
	}
	wg.Wait()
}

func tryChannel2() {
	ch := make(chan int)
	for j := 0; j<5; j++ {
		wg.Add(2)
		go func(ch <- chan int) {
			i := <- ch
			fmt.Println(i)
			wg.Done()
		}(ch)
		go func(ch chan <- int) {
			ch <- x
			x += 100
			wg.Done()
		}(ch)
	}
	wg.Wait()
}

func tryChannel3() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <- chan int) {
		for j := range ch {
			fmt.Println( j )
		}
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		for j := 0; j<5; j++ {
			ch <- x
			x += 100
		}
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

func tryChannel4() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <- chan int) {
		for  {
			if i, ok := <- ch; ok {
				fmt.Println( i )
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan <- int) {
		for j := 0; j<5; j++ {
			ch <- x
			x += 100
		}
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}


func main () {
	i := 4
	switch i {
	case 1: tryChannel1()
	case 2: tryChannel2()
	case 3: tryChannel3()
	case 4: tryChannel4()
	default: return
	}
}
