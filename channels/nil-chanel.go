package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int // チャネルをnilで初期化
	go func() {
		ch = make(chan int)
		time.Sleep(2 * time.Second) // 1秒後スリープ
		ch <- 1
	}()
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println(v)
			return
		default:
			fmt.Println("チャネルはまだnilです")
		}
	}
}
