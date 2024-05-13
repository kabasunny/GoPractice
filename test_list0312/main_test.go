package main

import (
	"testing"
	"time"
)

//以上追加

func main() {
	a := 100
	go func() { a += 50 }()
	go func() { a -= 50 }()

	time.Sleep(100 * time.Millisecond)
}

// 以下追加
func TestMain(t *testing.T) {
	main()
}
