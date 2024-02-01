package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestDeadlock(t *testing.T) {
	var state int32
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			mu.Lock()
			state += int32(i)
			mu.Unlock()
		}(i)
	}
}
