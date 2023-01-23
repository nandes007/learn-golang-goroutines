package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var loker = sync.Mutex{}
var cound = sync.NewCond(&loker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cound.L.Lock()
	cound.Wait()
	fmt.Println("Done", value)
	cound.L.Unlock()
}

func TestCound(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(1 * time.Second)
	//		cound.Signal()
	//	}
	//}()

	go func() {
		time.Sleep(1 * time.Second)
		cound.Broadcast()
	}()

	group.Wait()
}
