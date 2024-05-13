package main

import (
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestGoRoutineLeak(t *testing.T) {
	defer goleak.VerifyNone(t)

	go func() {
		time.Sleep(1 * time.Second)
	}()
}

/*
func main(){
	go time.Sleep(1 * time.Second)
}

func TestMain(m *testing.M){
	goleak.VerifyTestMain(m)
}

func Test_main(t *testing.T){
	main()
}

*/
