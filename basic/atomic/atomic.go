// atomic Increment and Get
package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义结构
type AtomicInt struct {
	value int
	lock  sync.Mutex
}

// 增加
func (a *AtomicInt) Increment() {
	fmt.Println("safe Increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

// 获取
func (a *AtomicInt) Get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a AtomicInt
	a.Increment()
	go func() {
		a.Increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.Get())
}
