package main

import (
	"go.uber.org/atomic"
	"sync"
)

func main() {

	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

	lock.Lock()
	for mailbox == 1 {
		sendCond.Wait()
	}
	mailbox = 1
	lock.Unlock()
	recvCond.Signal()
}