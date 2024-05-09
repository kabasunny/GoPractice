package main

import (
	"fmt"
	"sync"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func print(x int, wg *sync.WaitGroup, y string) {
	fmt.Println(x, y)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(12)

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	d := make(chan int)
	e := make(chan int)
	f := make(chan int)
	g := make(chan int)
	h := make(chan int)

	go sum(s[:len(s)/2], c) // -9, 4, 0 を渡す
	go sum(s[:len(s)/2], d) // -9, 4, 0 を渡す
	go sum(s[:len(s)/2], e) // -9, 4, 0 を渡す
	go sum(s[:len(s)/2], f) // -9, 4, 0 を渡す
	go sum(s[:len(s)/2], g) // -9, 4, 0 を渡す
	go sum(s[:len(s)/2], h) // -9, 4, 0 を渡す

	go sum(s[len(s)/2:], c) // 7, 2, 8 を渡す
	go sum(s[len(s)/2:], d) // 7, 2, 8 を渡す
	go sum(s[len(s)/2:], e) // 7, 2, 8 を渡す
	go sum(s[len(s)/2:], f) // 7, 2, 8 を渡す
	go sum(s[len(s)/2:], g) // 7, 2, 8 を渡す
	go sum(s[len(s)/2:], h) // 7, 2, 8 を渡す

	go func() {
		x := <-c
		go print(x, &wg, "c")
	}()
	go func() {
		x := <-d
		go print(x, &wg, "d")
	}()
	go func() {
		x := <-e
		go print(x, &wg, "e")
	}()
	go func() {
		x := <-f
		go print(x, &wg, "d")
	}()
	go func() {
		x := <-g
		go print(x, &wg, "f")
	}()
	go func() {
		x := <-h
		go print(x, &wg, "f")
	}()
	go func() {
		x := <-c
		go print(x, &wg, "c")
	}()
	go func() {
		x := <-d
		go print(x, &wg, "d")
	}()
	go func() {
		x := <-e
		go print(x, &wg, "e")
	}()
	go func() {
		x := <-f
		go print(x, &wg, "d")
	}()
	go func() {
		x := <-g
		go print(x, &wg, "f")
	}()
	go func() {
		x := <-h
		go print(x, &wg, "f")
	}()

	wg.Wait()
}

/* A Tour of Go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
*/
