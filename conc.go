package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

//Do parallelism
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func boo() {
	for i := 0; i <= 30; i++ {
		fmt.Println("BOO", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func foo() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Println("FOO", string(i))
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go boo()
	go foo()
	wg.Wait()
}
