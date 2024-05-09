package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second) // 二秒
		ch1 <- 1
	}()

	go func() {
		time.Sleep(3 * time.Second) //一秒
		ch2 <- 2
	}()

	for i := 0; i < 2; i++ {
		select { //コードの記載順序に関係なく、先に到達した方が実行される
		case x := <-ch1:
			fmt.Println(x)
		case y := <-ch2:
			fmt.Println(y)
		}
	}
}

/* A Tour of Goのselect
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for { //3.無限ループ
		select {
		case c <- x: //4.ここで1に値が送られる、10回繰り返す
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //1.ここでいったん止まり、2が実行され、10回送られてきた後、forを抜ける
		}
		quit <- 0 //5.2の呼び出しを終了させる
	}()
	fibonacci(c, quit) //2.呼び出す
}

*/
