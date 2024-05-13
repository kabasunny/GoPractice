package main

import (
	"testing"
	"time"

	"go.uber.org/goleak"
)

func main() {
	go time.Sleep(1 * time.Second)
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func Test_main(t *testing.T) {
	main()
}
