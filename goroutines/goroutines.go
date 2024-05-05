package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done() // 終了したらDoneを呼び出すことで、Addで渡した数字が‐1される
}

func main() {
	var wg sync.WaitGroup // WaitGroupを作成
	wg.Add(2)             // 待つゴルーチンの数を2つ増やす
	go say("world", &wg)  // ゴルーチンを立ち上げ、WaitGroupを渡す
	say("hello", &wg)     // メインゴルーチン？
	wg.Wait()             // 全てのゴルーチンが終了するのを待つ
}

/* A Tour of Goのコード
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
*/
