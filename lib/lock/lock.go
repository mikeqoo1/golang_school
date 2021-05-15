package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	balance int
	mux     sync.Mutex //互斥鎖
	//mux sync.RWMutex //讀寫鎖
}

func (b *Bank) 存錢(amount int) {
	b.mux.Lock()
	time.Sleep(time.Second) // spend 1 second
	b.balance += amount
	b.mux.Unlock()
}

func (b *Bank) Balance() (balance int) {
	b.mux.Lock()
	//b.mux.RLock()
	time.Sleep(time.Second) // spend 1 second
	balance = b.balance
	b.mux.Unlock()
	//b.mux.RUnlock()
	return
}

func main() {
	var wg sync.WaitGroup
	b := &Bank{}
	now := time.Now()

	n := 5
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			b.存錢(1000)
			fmt.Printf("Write: deposit amonut: %v\n", 1000)
			wg.Done()
		}()
	}
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			fmt.Printf("Read: balance: %v\n", b.Balance())
			wg.Done()
		}()
	}

	wg.Wait()

	duration := time.Since(now)
	fmt.Println("執行速度:", duration)
}

//-race 參數是 go 的 Race Detector，內建整合工具，可以輕鬆檢查出是否有 race condition
//go run -race lock.go
