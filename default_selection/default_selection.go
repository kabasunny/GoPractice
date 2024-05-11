package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Millisecond) //なぜかデフォルトが作動した
	//tick := time.Tick(80 * time.Millisecond)  //タイミングの裂け目
	//tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	c := 0 //forが何回実行されるのか気になって
	for {  //無限ループ
		select {
		case <-tick: //
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond) //ここをコメントアウトすると、デフォルト独壇場
		}
		c++
		fmt.Println(c) //forの速度...私のPCよりA Tour of Goのサーバー方がはるかに速い...
	}
}
