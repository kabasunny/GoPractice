package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for {
		fmt.Printf("Worker %d received data: %d\n", id, <-ch)
	}
}

func main() {
	ch := make(chan int, 2)

	// 2つのワーカーゴルーチンを起動
	for i := 0; i < 2; i++ {
		go worker(i, ch)
	}

	// チャネルにデータを送信
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

/*
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
*/
