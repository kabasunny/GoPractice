package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	sem := semaphore.NewWeighted(5)

	go do(sem, func() { time.Sleep(1 * time.Second) }, 1)
	go do(sem, func() { time.Sleep(1 * time.Second) }, 2)
	go do(sem, func() { time.Sleep(1 * time.Second) }, 3)
	go do(sem, func() { time.Sleep(1 * time.Second) }, 4)
	go do(sem, func() { time.Sleep(1 * time.Second) }, 5)
	// 	go do(sem, func() { time.Sleep(1 * time.Second) }, 6) エラー

	time.Sleep(6 * time.Second)
	fmt.Println("最高！")
}

func do(sem *semaphore.Weighted, f func(), w int64) { // ここの引数について
	if err := sem.Acquire(context.Background(), w); err != nil {
		log.Println(err)
		return
	}
	defer sem.Release(w)

	switch w {
	case 1:
		log.Printf("It will definitely come true!")
	case 2:
		log.Printf("That works!")
	case 3:
		log.Printf("I did it!")
	case 4:
		log.Printf("Thank you for this world!")
	case 5:
		log.Printf("Make the world better!")
	case 6:
		log.Printf("Do not lose!")
	default:
		log.Printf("How happy I am!")
	}

	f()
}
